package v1

import (
	"ecode/controllers"

	"github.com/gin-gonic/gin"
)

// Login 登录接口
func Login(c *gin.Context) {
	controllers.Login(c)
}

// UpdateToken 更新token
func UpdateToken(c *gin.Context) {
	controllers.UpdateToken(c)
}
