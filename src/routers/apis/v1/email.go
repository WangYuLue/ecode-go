package v1

import (
	C "ecode/controllers"
	M "ecode/utils/message"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ConfirmEmail 邮箱激活用户
func ConfirmEmail(c *gin.Context) {
	C.ConfirmEmail(c)
}

// SendConfirmEmail 发送激活用户邮件
func SendConfirmEmail(c *gin.Context) {
	if err := C.SendConfirmEmail(c); err != nil {
		M.HandelError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "邮件发送成功",
	})
}

// ResetPassword 通过邮箱重新设置密码
func ResetPassword(c *gin.Context) {
	if err := C.ResetPassword(c); err != nil {
		M.HandelError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "密码修改成功",
	})
}

// SendResetPasswordEmail 发送找回密码邮件
func SendResetPasswordEmail(c *gin.Context) {
	if err := C.SendResetPasswordEmail(c); err != nil {
		M.HandelError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "邮件发送成功",
	})
}
