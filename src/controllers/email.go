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
	redis.DB.HSet(redisKeys.EmailConfirm, strconv.Itoa(user.ID), uuidStr)
	confirmURL := config.EmailConfirm.BaseURL + "/v1/email/" + strconv.Itoa(user.ID) + "/confirm-email/" + uuidStr
	emailData := models.Mail{Name: user.Name, URL: confirmURL}
	emailTemplete, err := email.GenEmailHTML(emailData)
	if err != nil {
		log.Println("邮件模版生成异常")
	} else {
		// 发送用户激活邮件
		go email.SendEmailByAdmin(config.EmailConfirm.Title, emailTemplete, user.Email)
	}
}
