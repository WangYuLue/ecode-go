package email

import (
	"bytes"
	"ecode/config"
	"ecode/models"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"
)

// SendEmail 发送邮件
func SendEmail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var contentType string
	if mailtype == "html" {
		contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	message := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, message)
	return err
}

// SendEmailByAdmin 管理员发送邮件
func SendEmailByAdmin(title, html, to string) {
	user := config.Email.User
	password := config.Email.Password
	host := config.Email.Host

	fmt.Println("邮件发送中...")
	err := SendEmail(user, password, host, to, title, html, "html")
	if err != nil {
		fmt.Println("邮件发送失败")
		fmt.Println(err)
	} else {
		fmt.Println("邮件发送成功")
	}
}

// GenUserConfirmHTML 生成激活用户模版文件
func GenUserConfirmHTML(data models.Mail) (string, error) {
	var doc bytes.Buffer
	var err error
	var t *template.Template
	t, err = template.ParseFiles("templates/user-confirm.html")
	if err != nil {
		return "", err
	}
	err = t.Execute(&doc, data)
	if err != nil {
		return "", err
	}
	return doc.String(), nil
}

// GenResetPasswordHTML 生成激活用户模版文件
func GenResetPasswordHTML(data models.Mail) (string, error) {
	var doc bytes.Buffer
	var err error
	var t *template.Template
	t, err = template.ParseFiles("templates/reset-password.html")
	if err != nil {
		return "", err
	}
	err = t.Execute(&doc, data)
	if err != nil {
		return "", err
	}
	return doc.String(), nil
}
