package v1

import (
	C "ecode/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ConfirmEmail 邮箱激活用户
func ConfirmEmail(c *gin.Context) {
	C.ConfirmEmail(c)
}

// SendConfirmEmail 发送激活用户邮件
func SendConfirmEmail(c *gin.Context) {
	err := C.SendConfirmEmail(c)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "邮件发送成功",
		})
	}
}

// ResetPassword 通过邮箱重新设置密码
func ResetPassword(c *gin.Context) {
	err := C.ResetPassword(c)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "密码修改成功",
		})
	}
}

// SendResetPasswordEmail 发送找回密码邮件
func SendResetPasswordEmail(c *gin.Context) {
	err := C.SendResetPasswordEmail(c)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "邮件发送成功",
		})
	}
}
