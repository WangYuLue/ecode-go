package card

import (
	"net/http"

	C "ecode/controllers/card"

	"github.com/gin-gonic/gin"
)

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
