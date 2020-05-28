package message

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Errno 定义错误码
type Errno struct {
	Code    int
	Message string
}

func (err *Errno) Error() string {
	return err.Message
}

// ErrMsg 定义错误
type ErrMsg struct {
	Code    int    // 错误码
	Message string // 展示给用户看的
	Errord  error  // 保存内部错误信息
}

func (err *ErrMsg) Error() string {
	return fmt.Sprintf("ErrMsg - code: %d, message: %s, error: %s", err.Code, err.Message, err.Errord)
}

// NewErrMsg -
func NewErrMsg(errno *Errno, err error) *ErrMsg {
	return &ErrMsg{
		Code:    errno.Code,
		Message: errno.Message,
		Errord:  err,
	}
}

// DecodeErr 解码错误, 获取 Code 和 Message
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch typed := err.(type) {
	case *ErrMsg:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return ErrSystem.Internal.Code, err.Error()
}

func handelError(c *gin.Context, code int, errno *Errno) {
	c.JSON(code, gin.H{
		"code":    errno.Code,
		"message": errno.Message,
	})
}

// HandelError 检查并处理 400 异常
func HandelError(c *gin.Context, errno *Errno) {
	handelError(c, http.StatusBadRequest, errno)
}

// HandelStatusUnauthorizedError 检查并处理 401 异常
func HandelStatusUnauthorizedError(c *gin.Context, errno *Errno) {
	handelError(c, http.StatusUnauthorized, errno)
}
