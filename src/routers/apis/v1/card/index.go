package card

import (
	"net/http"
	"strconv"

	C "ecode/controllers/card"
	"ecode/models"
	"ecode/utils/message"

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
	id, err := strconv.Atoi(c.Param("cardid"))
	if err != nil {
		message.HandelError(c, message.ErrCard.IDIllegal)
		return
	}
	data, err := models.GetPublicCardByID(id)
	if err != nil {
		message.HandelError(c, message.ErrCard.NotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
