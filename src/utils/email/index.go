package email

import (
	"ecode/config"
	"ecode/databases/redis"
	redisKeys "ecode/databases/redis/keys"
	"ecode/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
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

// SendUserConfirmEmail 向刚注册的用户发送激活邮件
func SendUserConfirmEmail(user models.User) {
	uuidStr := uuid.Must(uuid.NewV4()).String()
	redis.DB.HSet(redisKeys.EmailConfirmUser, strconv.Itoa(user.ID), uuidStr)
	confirmURL := config.BaseURL + "/v1/confirm-email/" + strconv.Itoa(user.ID) + "/" + uuidStr
	data := models.Mail{Name: user.Name, URL: confirmURL}
	emailTemplete, err := GenUserConfirmHTML(data)
	if err != nil {
		log.Println("邮件模版生成异常")
	} else {
		// 发送用户激活邮件
		go SendEmailByAdmin(config.EmailConfirmUser.Title, emailTemplete, user.Email)
	}
}

// SendResetPasswordEmail 发送重新设置密码邮件
func SendResetPasswordEmail(user models.User) {
	uuidStr := uuid.Must(uuid.NewV4()).String()
	redis.DB.HSet(redisKeys.EmailResetPassword, strconv.Itoa(user.ID), uuidStr)
	resetPasswordURL := config.EmailResetPassword.ResetURL + "?id=" + strconv.Itoa(user.ID) + "&uuid=" + uuidStr
	data := models.Mail{Name: user.Name, URL: resetPasswordURL}
	emailTemplete, err := GenResetPasswordHTML(data)
	if err != nil {
		log.Println("重置密码模版生成异常")
	} else {
		// 发送用户激活邮件
		go SendEmailByAdmin(config.EmailConfirmUser.Title, emailTemplete, user.Email)
	}
}
