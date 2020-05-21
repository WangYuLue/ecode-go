package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type statusFailMessage struct {
	Add string
	Del string
	Mod string
	Get string
}

type statusIllegalMessage struct {
	ID   string
	Data string
}

type statusNoneMessage struct {
	User string
	Card string
}

type statusBadMessage struct {
	Fail    statusFailMessage
	Illegal statusIllegalMessage
	None    statusNoneMessage
}

// StatusBadMessage 请求异常返回的信息
var StatusBadMessage = statusBadMessage{
	Fail: statusFailMessage{
		Add: "添加失败",
		Del: "删除失败",
		Mod: "修改失败",
		Get: "查询失败",
	},
	Illegal: statusIllegalMessage{
		ID:   "ID不合法",
		Data: "数据不合法",
	},
	None: statusNoneMessage{
		User: "用户不存在",
		Card: "卡片不存在",
	},
}

// HandelError 检查并处理异常
func HandelError(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"data": errorMessage,
	})
}
