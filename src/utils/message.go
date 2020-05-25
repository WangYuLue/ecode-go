package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SStatusFailMessage 失败提示
type SStatusFailMessage struct {
	Add string
	Del string
	Mod string
	Get string
}

// SStatusIllegalMessage 非法提示
type SStatusIllegalMessage struct {
	ID   string
	Data string
}

// SStatusNoneMessage 不存在提示
type SStatusNoneMessage struct {
	User string
	Card string
}

// SStatusBadMessage 请求失败提示汇总
type SStatusBadMessage struct {
	Fail    SStatusFailMessage
	Illegal SStatusIllegalMessage
	None    SStatusNoneMessage
}

// StatusBadMessage 请求异常返回的信息
var StatusBadMessage = SStatusBadMessage{
	Fail: SStatusFailMessage{
		Add: "添加失败",
		Del: "删除失败",
		Mod: "修改失败",
		Get: "查询失败",
	},
	Illegal: SStatusIllegalMessage{
		ID:   "ID不合法",
		Data: "数据不合法",
	},
	None: SStatusNoneMessage{
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
