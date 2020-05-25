package email

import (
	"ecode/config"
	"fmt"
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
	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}

// SendEmailByAdmin 管理员发送邮件
func SendEmailByAdmin() {
	user := config.Email.User
	password := config.Email.Password
	host := config.Email.Host
	to := "1206180097@qq.com"
	// to := "wangyulue@gmail.com"

	subject := "使用Golang发送邮件"

	body := `
		<html>
		<body>
		<h3>
		"Test send to email"
		</h3>
		</body>
		</html>
		`
	fmt.Println("邮件发送中...")
	err := SendEmail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("邮件发送失败")
		fmt.Println(err)
	} else {
		fmt.Println("邮件发送成功")
	}
}
