package v1

import (
	"log"
	"net/http"
	"strconv"

	"ecode/controllers"
	"ecode/models"
	"ecode/utils"
	"ecode/utils/md5"

	"github.com/gin-gonic/gin"
)

// AddUserAPI 添加 card
func AddUserAPI(c *gin.Context) {
	// name := c.Request.FormValue("name")
	nameStr := c.PostForm("name")
	emailStr := c.PostForm("email")
	passwordStr := c.PostForm("password")
	log.Println(nameStr, passwordStr, emailStr)
	if nameStr == "" || passwordStr == "" || emailStr == "" {
		utils.HandelError(c, utils.StatusBadMessage.Illegal.Data)
		return
	}
	if users, err := models.GetUserByName(nameStr); err != nil || len(users) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": "用户名已存在",
		})
		return
	}
	if users, err := models.GetUserByEmail(emailStr); err != nil || len(users) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": "邮箱已注册",
		})
		return
	}
	passwordStr = md5.Md5(passwordStr)
	p := &models.User{
		Name:     nameStr,
		Email:    emailStr,
		Password: passwordStr,
	}
	user, err := models.AddUser(p)
	if err != nil {
		utils.HandelError(c, utils.StatusBadMessage.Fail.Add)
		return
	}
	controllers.SendUserConfirmEmail(user)
	c.JSON(http.StatusOK, gin.H{
		"data": "用户注册成功",
	})
}

// GetUsersAPI 获取所有 card
func GetUsersAPI(c *gin.Context) {
	data, err := models.GetUsers()
	if err != nil {
		utils.HandelError(c, utils.StatusBadMessage.Fail.Get)
		return
	}
	// out, _ := StructPick(data, models.UserMap)
	// c.JSON(http.StatusOK, gin.H{
	// 	"data": out,
	// })
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// GetUserAPI 根据 ID 获取 card
func GetUserAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		utils.HandelError(c, utils.StatusBadMessage.Illegal.ID)
		return
	}
	data, err := models.GetUserByID(id)
	if err != nil {
		utils.HandelError(c, utils.StatusBadMessage.None.User)
		return
	}
	out, _ := StructPick(data, models.UserMap)
	c.JSON(http.StatusOK, gin.H{
		"data": out,
	})
}

// GetCardsByUserID 根据 ID 获取 card
func GetCardsByUserID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		utils.HandelError(c, utils.StatusBadMessage.Illegal.ID)
		return
	}
	data, err := models.GetCardsByUserID(id)
	if err != nil {
		utils.HandelError(c, utils.StatusBadMessage.None.User)
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
		utils.HandelError(c, utils.StatusBadMessage.Illegal.ID)
		return
	}
	_, err = models.GetUserByID(id)
	if err != nil {
		utils.HandelError(c, utils.StatusBadMessage.None.User)
		return
	}
	err = models.ModUserByID(id, name)
	if err != nil {
		utils.HandelError(c, utils.StatusBadMessage.Fail.Mod)
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
		utils.HandelError(c, utils.StatusBadMessage.Illegal.ID)
		return
	}
	_, err = models.GetUserByID(id)
	if err != nil {
		utils.HandelError(c, utils.StatusBadMessage.None.User)
		return
	}
	err = models.DelUserByID(id)
	if err != nil {
		utils.HandelError(c, utils.StatusBadMessage.Fail.Del)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "删除成功",
	})
}
