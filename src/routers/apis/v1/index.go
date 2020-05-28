package v1

import (
	"ecode/models"
	"ecode/utils/lodash"
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexAPI 测试服务器是否启动
func IndexAPI(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// MapLowFirstCase map 所有 key 首字母小写
func MapLowFirstCase(in models.H) models.H {
	out := make(models.H)
	for k, v := range in {
		out[lodash.ToLowerFirstCase(k)] = v
	}
	return out
}
