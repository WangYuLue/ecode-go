package user

import (
	C "ecode/controllers/user"
	M "ecode/utils/message"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register 注册用户
func Register(c *gin.Context) {
	if err := C.Register(c); err != nil {
		M.HandelError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "用户注册成功",
	})
}

// GetUsers 获取所有 user
func GetUsers(c *gin.Context) {
	data, err := C.GetUsers(c)
	if err != nil {
		M.HandelError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// GetUser 根据 ID 获取 user
func GetUser(c *gin.Context) {
	data, err := C.GetUser(c)
	if err != nil {
		M.HandelError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// GetCardsByUserID 根据 ID 获取 user
func GetCardsByUserID(c *gin.Context) {
	data, err := C.GetCardsByUserID(c)
	if err != nil {
		M.HandelError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// ModUser 修改用户
func ModUser(c *gin.Context) {
	if err := C.ModUser(c); err != nil {
		M.HandelError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "修改成功",
	})
}

// DelUser 删除用户
func DelUser(c *gin.Context) {
	if err := C.DelUser(c); err != nil {
		M.HandelError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}
