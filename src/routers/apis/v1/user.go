package v1

import (
	"net/http"
	"strconv"

	"ecode/models"
	"ecode/utils"
	"ecode/utils/email"

	"github.com/gin-gonic/gin"
)

// AddUserAPI 添加 card
func AddUserAPI(c *gin.Context) {
	// name := c.Request.FormValue("name")
	// name := c.PostForm("name")
	var u models.User
	if c.ShouldBind(&u) != nil {
		utils.HandelError(c, utils.StatusBadMessage.Illegal.Data)
		return
	}
	if models.AddUser(&u) != nil {
		utils.HandelError(c, utils.StatusBadMessage.Fail.Add)
		return
	}
	emailTemplete := email.UserConfirmTemplete("", "", "")

	email.SendEmailByAdmin(email.UserConfirmTitle(), emailTemplete, "wangyulue@gmail.com")
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
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// GetUserAPI 根据 ID 获取 card
func GetUserAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.HandelError(c, utils.StatusBadMessage.Illegal.ID)
		return
	}
	data, err := models.GetUserByID(id)
	if err != nil {
		utils.HandelError(c, utils.StatusBadMessage.None.User)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// GetCardsByUserID 根据 ID 获取 card
func GetCardsByUserID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
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
	id, err := strconv.Atoi(c.Param("id"))
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
	id, err := strconv.Atoi(c.Param("id"))
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
