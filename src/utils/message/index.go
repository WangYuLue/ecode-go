package message

import (
	"ecode/utils/log"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrMsg 定义错误
type ErrMsg struct {
	Code    int    // 错误码
	Message string // 展示给用户看的
	Detail  error  // 保存内部错误信息
}

func (err ErrMsg) Error() string {
	if err.Detail == nil {
		return fmt.Sprintf("ErrMsg - code: %d, message: %s", err.Code, err.Message)
	}
	return fmt.Sprintf("ErrMsg - code: %d, message: %s, detail: %s", err.Code, err.Message, err.Detail)
}

// NewErrMsg -
func NewErrMsg(errmsg ErrMsg, detail error) ErrMsg {
	return ErrMsg{
		Code:    errmsg.Code,
		Message: errmsg.Message,
		Detail:  detail,
	}
}

// DecodeErr 解码错误, 获取 Code 和 Message
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch typed := err.(type) {
	case ErrMsg:
		return typed.Code, typed.Message
	default:
	}

	return ErrSystem.Internal.Code, ErrSystem.Internal.Message
}

func handelError(c *gin.Context, httpCode int, err error) {
	log.Error(c.ClientIP(), c.Request.URL, err.Error())
	code, message := DecodeErr(err)
	c.JSON(httpCode, gin.H{
		"code":    code,
		"message": message,
	})
}

// HandelError 检查并处理 400 异常
func HandelError(c *gin.Context, err error) {
	fmt.Println(err)
	handelError(c, http.StatusBadRequest, err)
}

// HandelStatusUnauthorizedError 检查并处理 401 异常
func HandelStatusUnauthorizedError(c *gin.Context, err error) {
	handelError(c, http.StatusUnauthorized, err)
}
