package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexAPI 测试服务器是否启动
func IndexAPI(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
