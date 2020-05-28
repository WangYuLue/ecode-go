package v1

import (
	"log"
	"net/http"
	"strconv"

	"ecode/controllers"
	"ecode/models"
	"ecode/utils/md5"
	"ecode/utils/message"

	"github.com/gin-gonic/gin"
)

// RegisterAPI 注册用户
func RegisterAPI(c *gin.Context) {
	// name := c.Request.FormValue("name")
	nameStr := c.PostForm("name")
	emailStr := c.PostForm("email")
	passwordStr := c.PostForm("password")
	log.Println(nameStr, passwordStr, emailStr)
	if nameStr == "" || passwordStr == "" || emailStr == "" {
		message.HandelError(c, message.ErrHTTPData.BindFail)
		return
	}
	if user, _ := models.GetUserByName(nameStr); user.Name != "" {
		message.HandelError(c, message.ErrUser.NameExist)
		return
	}
	if user, _ := models.GetUserByEmail(emailStr); user.Email != "" {
		message.HandelError(c, message.ErrUser.EmailExist)
		return
	}
	passwordStr = md5.Md5(passwordStr)
	p := &models.UserORM{
		Name:     nameStr,
		Email:    emailStr,
		Password: passwordStr,
	}
	err := models.AddUser(p)
	if err != nil {
		message.HandelError(c, message.ErrUser.ADDFail)
		return
	}
	controllers.SendUserConfirmEmail(models.User{
		Name:  nameStr,
		Email: emailStr,
	})
	c.JSON(http.StatusOK, gin.H{
		"data": "用户注册成功",
	})
}

// GetUsersAPI 获取所有 user
func GetUsersAPI(c *gin.Context) {
	data, err := models.GetUsers()
	if err != nil {
		message.HandelError(c, message.ErrUser.NotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// GetUserAPI 根据 ID 获取 user
func GetUserAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return
	}
	data, err := models.GetUserByID(id)
	if err != nil {
		message.HandelError(c, message.ErrUser.NotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// GetCardsByUserID 根据 ID 获取 user
func GetCardsByUserID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return
	}
	data, err := models.GetCardsByUserID(id)
	if err != nil {
		message.HandelError(c, message.ErrUser.NotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// ModUserAPI 修改用户
func ModUserAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userid"))
	name := c.PostForm("name")
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return
	}
	_, err = models.GetUserByID(id)
	if err != nil {
		message.HandelError(c, message.ErrUser.NotFound)
		return
	}
	err = models.ModUserByID(id, name)
	if err != nil {
		message.HandelError(c, message.ErrUser.ModFail)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "修改成功",
	})
}

// DelUserAPI 删除用户
func DelUserAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return
	}
	_, err = models.GetUserByID(id)
	if err != nil {
		message.HandelError(c, message.ErrUser.NotFound)
		return
	}
	err = models.DelUserByID(id)
	if err != nil {
		message.HandelError(c, message.ErrUser.DelFail)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "删除成功",
	})
}
