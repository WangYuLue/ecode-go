package user

import (
	"ecode/models"
	"ecode/utils/email"
	"ecode/utils/md5"
	M "ecode/utils/message"
	"strconv"

	"github.com/gin-gonic/gin"
)

// // ErrorDefault -
// var ErrorDefault = errors.New("")

// // CardDefault -
// var CardDefault = models.Card{}

// // CategoryDefault -
// var CategoryDefault = models.Category{}

// // TagDefault -
// var TagDefault = models.Tag{}

// // UserDefault -
// var UserDefault = models.User{}

// Register 注册用户
func Register(c *gin.Context) error {
	// name := c.Request.FormValue("name")
	nameStr := c.PostForm("name")
	emailStr := c.PostForm("email")
	passwordStr := c.PostForm("password")
	if nameStr == "" || passwordStr == "" || emailStr == "" {
		return M.ErrHTTPData.BindFail
	}
	if user, _ := models.GetUserByName(nameStr); user.Name != "" {
		return M.ErrUser.NameExist
	}
	if user, _ := models.GetUserByEmail(emailStr); user.Email != "" {
		return M.ErrUser.EmailExist
	}
	passwordStr = md5.Md5(passwordStr)
	p := &models.UserORM{
		Name:     nameStr,
		Email:    emailStr,
		Password: passwordStr,
	}
	user, err := models.AddUser(p)
	if err != nil {
		return M.NewErrMsg(M.ErrUser.AddFail, err)
	}
	email.SendUserConfirmEmail(models.User{
		ID:    user.ID,
		Name:  nameStr,
		Email: emailStr,
	})
	return nil
}

// GetUsers 获取所有 user
func GetUsers(c *gin.Context) ([]models.User, error) {
	data, err := models.GetUsers()
	if err != nil {
		return nil, M.NewErrMsg(M.ErrUser.NotFound, err)
	}
	return data, nil
}

// GetUser 根据 ID 获取 user
func GetUser(c *gin.Context) (models.User, error) {
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return models.User{}, M.ErrUser.IDIllegal
	}
	data, err := models.GetUserByID(id)
	if err != nil {
		return models.User{}, M.NewErrMsg(M.ErrUser.NotFound, err)
	}
	return data, nil
}

// GetCardsByUserID 根据 ID 获取 user
func GetCardsByUserID(c *gin.Context) ([]models.CardORM, error) {
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return nil, M.ErrUser.IDIllegal
	}
	data, err := models.GetCardsByUserID(id)
	if err != nil {
		return nil, M.NewErrMsg(M.ErrUser.NotFound, err)
	}
	return data, nil
}

// ModUser 修改用户
func ModUser(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("userid"))
	name := c.PostForm("name")
	if err != nil {
		return M.ErrUser.IDIllegal
	}
	if _, err = models.GetUserByID(id); err != nil {
		return M.NewErrMsg(M.ErrUser.NotFound, err)
	}
	if err = models.ModUserByID(id, models.UserORM{Name: name}); err != nil {
		return M.NewErrMsg(M.ErrUser.ModFail, err)
	}
	return nil
}

// DelUser 删除用户
func DelUser(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return M.ErrUser.IDIllegal
	}
	if _, err = models.GetUserByID(id); err != nil {
		return M.NewErrMsg(M.ErrUser.NotFound, err)
	}
	if err = models.DelUserByID(id); err != nil {
		return M.NewErrMsg(M.ErrUser.DelFail, err)
	}
	return nil
}
