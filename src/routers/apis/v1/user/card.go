package user

import (
	C "ecode/controllers/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddCard 添加卡片
func AddCard(c *gin.Context) {
	err := C.AddCard(c)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "发布卡片成功",
		})
	}
}

// GetCards 获取所有卡片
func GetCards(c *gin.Context) {
	data, err := C.GetCards(c)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

// GetCard 根据 ID 获取卡片
func GetCard(c *gin.Context) {
	data, err := C.GetCard(c)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

// ModCard 修改卡片
func ModCard(c *gin.Context) {
	err := C.ModCard(c)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "修改成功",
		})
	}
}

// DelCard 删除卡片
func DelCard(c *gin.Context) {
	err := C.DelCard(c)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "删除成功",
		})
	}
}
