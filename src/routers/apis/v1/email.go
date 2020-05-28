package v1

import (
	"ecode/config"
	"ecode/controllers"
	"ecode/databases/redis"
	redisKeys "ecode/databases/redis/keys"
	"ecode/models"
	"ecode/utils/message"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ConfirmEmail 邮箱激活用户
func ConfirmEmail(c *gin.Context) {
	id := c.Param("userid")
	uuid1 := c.Param("uuid")
	if id == "" || uuid1 == "" {
		// 重定向到登录失败页面
		c.Redirect(http.StatusMovedPermanently, config.EmailConfirm.FailURL+"?message=激活链接不合法")
		return
	}
	uuid2 := redis.DB.HGet(redisKeys.EmailConfirm, id).Val()
	if uuid1 != uuid2 {
		// 重定向到登录失败页面
		c.Redirect(http.StatusMovedPermanently, config.EmailConfirm.FailURL+"?message=验证失败")
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, config.EmailConfirm.FailURL+"?message=ID不合法")
		return
	}
	_, err = models.ActiveUser(idInt)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, config.EmailConfirm.FailURL+"?message=激活失败")
		return
	}
	// 重定向到登录成功页面
	c.Redirect(http.StatusMovedPermanently, config.EmailConfirm.SuccessURL)
}

// SendConfirmEmail 重新发送激活邮件
func SendConfirmEmail(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return
	}
	user, err := models.GetUserByID(idInt)
	if err != nil {
		message.HandelError(c, message.ErrUser.NotFound)
		return
	}
	controllers.SendUserConfirmEmail(user)
	c.JSON(http.StatusOK, gin.H{
		"data": "邮件发送成功",
	})
}
