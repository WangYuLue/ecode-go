package card

import (
	"strconv"

	"ecode/models"
	"ecode/utils/message"

	"github.com/gin-gonic/gin"
)

// GetCards 获取所有卡片
func GetCards(c *gin.Context) ([]models.Card, error) {
	cards, err := models.GetPublicCards()
	if err != nil {
		message.HandelError(c, message.ErrCard.NotFound)
		return nil, err
	}
	return cards, nil
}

// GetCard 根据 ID 获取卡片
func GetCard(c *gin.Context) (models.Card, error) {
	id, err := strconv.Atoi(c.Param("cardid"))
	if err != nil {
		message.HandelError(c, message.ErrCard.IDIllegal)
		return models.Card{}, err
	}
	card, err := models.GetPublicCardByID(id)
	if err != nil {
		message.HandelError(c, message.ErrCard.NotFound)
		return models.Card{}, err
	}
	return card, nil
}
