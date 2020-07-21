package user

import (
	C "ecode/controllers/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddTag 添加标签
func AddTag(c *gin.Context) {
	err := C.AddTag(c)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "发布标签成功",
		})
	}
}

// GetTags 获取所有标签
func GetTags(c *gin.Context) {
	data, err := C.GetTags(c)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

// GetTag 根据 ID 获取标签
func GetTag(c *gin.Context) {
	data, err := C.GetTag(c)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

// ModTag 修改标签
func ModTag(c *gin.Context) {
	err := C.ModTag(c)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "修改成功",
		})
	}
}

// DelTag 删除标签
func DelTag(c *gin.Context) {
	err := C.DelTag(c)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "删除成功",
		})
	}
}

// AddCardToTag -
func AddCardToTag(c *gin.Context) {
	// TODO:
	c.JSON(http.StatusOK, gin.H{
		"message": "添加成功",
	})
}

// RemoveCardToTag -
func RemoveCardToTag(c *gin.Context) {
	// TODO:
	c.JSON(http.StatusOK, gin.H{
		"message": "移除成功",
	})
}
