package user

import (
	C "ecode/controllers/user"
	M "ecode/utils/message"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddCategory 添加分类
func AddCategory(c *gin.Context) {
	if err := C.AddCategory(c); err != nil {
		M.HandelError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "发布分类成功",
	})
}

// GetCategorys 获取所有分类
func GetCategorys(c *gin.Context) {
	data, err := C.GetCategorys(c)
	if err != nil {
		M.HandelError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// GetCategory 根据 ID 获取分类
func GetCategory(c *gin.Context) {
	data, err := C.GetCategory(c)
	if err != nil {
		M.HandelError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// ModCategory 修改分类
func ModCategory(c *gin.Context) {
	if err := C.ModCategory(c); err != nil {
		M.HandelError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "修改成功",
	})
}

// DelCategory 删除分类
func DelCategory(c *gin.Context) {
	if err := C.DelCategory(c); err != nil {
		M.HandelError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

// AddCardToCategory -
func AddCardToCategory(c *gin.Context) {
	if err := C.AddCardToCategory(c); err != nil {
		M.HandelError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "添加成功",
	})
}

// RemoveCardToCategory -
func RemoveCardToCategory(c *gin.Context) {
	if err := C.RemoveCardToCategory(c); err != nil {
		M.HandelError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "移除成功",
	})
}
