package controllers

import (
	"ecode/config"
	"ecode/databases/redis"
	redisKeys "ecode/databases/redis/keys"
	"ecode/models"
	"ecode/utils/email"
	"log"
	"strconv"

	"github.com/gofrs/uuid"
)

// SendUserConfirmEmail 向刚注册的用户发送激活邮件
func SendUserConfirmEmail(user models.User) {
	uuidStr := uuid.Must(uuid.NewV4()).String()
	redis.DB.HSet(redisKeys.EmailConfirmUser, strconv.Itoa(user.ID), uuidStr)
	confirmURL := config.BaseURL + "/v1/confirm-email/" + strconv.Itoa(user.ID) + "/" + uuidStr
	data := models.Mail{Name: user.Name, URL: confirmURL}
	emailTemplete, err := email.GenUserConfirmHTML(data)
	if err != nil {
		log.Println("邮件模版生成异常")
	} else {
		// 发送用户激活邮件
		go email.SendEmailByAdmin(config.EmailConfirmUser.Title, emailTemplete, user.Email)
	}
}

// SendResetPasswordEmail 发送重新设置密码邮件
func SendResetPasswordEmail(user models.User) {
	uuidStr := uuid.Must(uuid.NewV4()).String()
	redis.DB.HSet(redisKeys.EmailResetPassword, strconv.Itoa(user.ID), uuidStr)
	resetPasswordURL := config.EmailResetPassword.ResetURL + "?id=" + strconv.Itoa(user.ID) + "&uuid=" + uuidStr
	data := models.Mail{Name: user.Name, URL: resetPasswordURL}
	emailTemplete, err := email.GenResetPasswordHTML(data)
	if err != nil {
		log.Println("重置密码模版生成异常")
	} else {
		// 发送用户激活邮件
		go email.SendEmailByAdmin(config.EmailConfirmUser.Title, emailTemplete, user.Email)
	}
}
