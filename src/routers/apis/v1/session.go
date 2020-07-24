package v1

import (
	C "ecode/controllers"
	M "ecode/utils/message"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login 登录接口
func Login(c *gin.Context) {
	data, err := C.Login(c)
	if err != nil {
		M.HandelError(c, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

// UpdateToken 更新token
func UpdateToken(c *gin.Context) {
	data, err := C.UpdateToken(c)
	if err != nil {
		M.HandelError(c, err)
	}
	c.JSON(http.StatusOK, data)
}
