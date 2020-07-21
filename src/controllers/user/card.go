package user

import (
	"strconv"

	"ecode/models"
	"ecode/utils/message"

	"github.com/gin-gonic/gin"
)

// AddCard 添加卡片
func AddCard(c *gin.Context) error {
	var u models.CardORM
	if c.ShouldBind(&u) != nil {
		message.HandelError(c, message.ErrHTTPData.BindFail)
		return ErrorDefault
	}
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return ErrorDefault
	}
	u.AutherID = userid
	if models.AddCard(&u) != nil {
		message.HandelError(c, message.ErrCard.AddFail)
		return ErrorDefault
	}
	return nil
}

// GetCards 获取所有卡片
func GetCards(c *gin.Context) ([]models.Card, error) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return nil, err
	}
	data, err := models.GetPrivateCards(userid)
	if err != nil {
		message.HandelError(c, message.ErrCard.NotFound)
		return nil, err
	}
	return data, nil
}

// GetCard 根据 ID 获取卡片
func GetCard(c *gin.Context) (models.Card, error) {

	id, err := strconv.Atoi(c.Param("cardid"))
	if err != nil {
		message.HandelError(c, message.ErrCard.IDIllegal)
		return CardDefault, err
	}
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return CardDefault, err
	}
	data, err := models.GetPrivateCardByID(userid, id)
	if err != nil {
		message.HandelError(c, message.ErrCard.NotFound)
		return CardDefault, err
	}
	return data, nil
}

// ModCard 修改卡片
func ModCard(c *gin.Context) error {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return ErrorDefault
	}
	id, err := strconv.Atoi(c.Param("cardid"))
	if err != nil {
		message.HandelError(c, message.ErrCard.IDIllegal)
		return ErrorDefault
	}
	question := c.PostForm("question")
	answer := c.PostForm("answer")

	_, err = models.GetPrivateCardByID(userid, id)
	if err != nil {
		message.HandelError(c, message.ErrCard.NotFound)
		return ErrorDefault
	}
	err = models.ModCardByID(id, question, answer)
	if err != nil {
		message.HandelError(c, message.ErrCard.ModFail)
		return ErrorDefault
	}
	return nil
}

// DelCard 删除卡片
func DelCard(c *gin.Context) error {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return ErrorDefault
	}
	id, err := strconv.Atoi(c.Param("cardid"))
	if err != nil {
		message.HandelError(c, message.ErrCard.IDIllegal)
		return ErrorDefault
	}
	_, err = models.GetPrivateCardByID(userid, id)
	if err != nil {
		message.HandelError(c, message.ErrCard.NotFound)
		return ErrorDefault
	}
	err = models.DelCardByID(id)
	if err != nil {
		message.HandelError(c, message.ErrCard.DelFail)
		return ErrorDefault
	}
	return nil
}
