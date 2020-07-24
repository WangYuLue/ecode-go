package controllers

import (
	"ecode/config"
	"ecode/databases/redis"
	redisKeys "ecode/databases/redis/keys"
	"ecode/models"
	"ecode/utils/email"
	"ecode/utils/md5"
	M "ecode/utils/message"
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
func SendConfirmEmail(c *gin.Context) error {
	idInt, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return M.ErrUser.IDIllegal
	}
	user, err := models.GetUserByID(idInt)
	if err != nil {
		return M.NewErrMsg(M.ErrUser.NotFound, err)
	}
	email.SendUserConfirmEmail(user)
	return nil
}

// ResetPassword 通过邮箱重新设置密码
func ResetPassword(c *gin.Context) error {
	id := c.PostForm("id")
	uuid1 := c.PostForm("uuid")
	password := c.PostForm("password")
	uuid2 := redis.DB.HGet(redisKeys.EmailResetPassword, id).Val()
	if uuid1 != uuid2 {
		return M.ErrUser.UUIDIllegal
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return M.ErrHTTPData.BindFail
	}
	password = md5.Md5(password)
	err = models.ModUserByID(idInt, models.UserORM{Password: password})
	if err != nil {
		return M.NewErrMsg(M.ErrUser.ModPasswordFail, err)
	}
	return nil
}

// SendResetPasswordEmail 发送找回密码邮件
func SendResetPasswordEmail(c *gin.Context) error {
	emailStr := c.Param("email")
	user, err := models.GetUserByEmail(emailStr)
	if err != nil {
		return M.ErrUser.NotFound
	}
	email.SendResetPasswordEmail(user)
	return nil
}
