package v1

import (
	"ecode/config"
	"ecode/controllers"
	"ecode/databases/redis"
	redisKeys "ecode/databases/redis/keys"
	"ecode/models"
	"ecode/utils/md5"
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
		c.Redirect(http.StatusMovedPermanently, config.EmailConfirmUser.FailURL+"?message=激活链接不合法")
		return
	}
	uuid2 := redis.DB.HGet(redisKeys.EmailConfirmUser, id).Val()
	if uuid1 != uuid2 {
		// 重定向到登录失败页面
		c.Redirect(http.StatusMovedPermanently, config.EmailConfirmUser.FailURL+"?message=验证失败")
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, config.EmailConfirmUser.FailURL+"?message=ID不合法")
		return
	}
	_, err = models.ActiveUser(idInt)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, config.EmailConfirmUser.FailURL+"?message=激活失败")
		return
	}
	// 重定向到登录成功页面
	c.Redirect(http.StatusMovedPermanently, config.EmailConfirmUser.SuccessURL)
}

// SendConfirmEmail 发送激活用户邮件
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
		"message": "邮件发送成功",
	})
}

// ResetPassword 通过邮箱重新设置密码
func ResetPassword(c *gin.Context) {
	id := c.PostForm("id")
	uuid1 := c.PostForm("uuid")
	password := c.PostForm("password")
	uuid2 := redis.DB.HGet(redisKeys.EmailResetPassword, id).Val()
	if uuid1 != uuid2 {
		message.HandelError(c, message.ErrUser.UUIDIllegal)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		message.HandelError(c, message.ErrHTTPData.BindFail)
		return
	}
	password = md5.Md5(password)
	err = models.ModUserByID(idInt, models.UserORM{Password: password})
	if err != nil {
		message.HandelError(c, message.ErrUser.ModPasswordFail)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "密码修改成功",
	})
}

// SendResetPasswordEmail 发送找回密码邮件
func SendResetPasswordEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := models.GetUserByEmail(email)
	if err != nil {
		message.HandelError(c, message.ErrUser.NotFound)
		return
	}
	controllers.SendResetPasswordEmail(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "邮件发送成功",
	})
}
