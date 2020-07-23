package card

import (
	"strconv"

	"ecode/models"
	M "ecode/utils/message"

	"github.com/gin-gonic/gin"
)

// GetCards 获取所有卡片
func GetCards(c *gin.Context) ([]models.Card, error) {
	cards, err := models.GetPublicCards()
	if err != nil {
		M.HandelError(c, M.NewErrMsg(M.ErrCard.NotFound, err))
		return nil, err
	}
	return cards, nil
}

// GetCard 根据 ID 获取卡片
func GetCard(c *gin.Context) (models.Card, error) {
	id, err := strconv.Atoi(c.Param("cardid"))
	if err != nil {
		M.HandelError(c, M.ErrCard.IDIllegal)
		return models.Card{}, err
	}
	card, err := models.GetPublicCardByID(id)
	if err != nil {
		M.HandelError(c, M.NewErrMsg(M.ErrCard.NotFound, err))
		return models.Card{}, err
	}
	return card, nil
}
