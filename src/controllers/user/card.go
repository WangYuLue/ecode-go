package user

import (
	"strconv"

	"ecode/models"
	M "ecode/utils/message"

	"github.com/gin-gonic/gin"
)

// AddCard 添加卡片
func AddCard(c *gin.Context) error {
	var u models.CardORM
	if c.ShouldBind(&u) != nil {
		return M.ErrHTTPData.BindFail
	}
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return M.ErrUser.IDIllegal
	}
	u.AutherID = userid
	if err = models.AddCard(&u); err != nil {
		return M.NewErrMsg(M.ErrCard.AddFail, err)
	}
	return nil
}

// GetCards 获取所有卡片
func GetCards(c *gin.Context) ([]models.Card, error) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return nil, M.ErrUser.IDIllegal
	}
	data, err := models.GetPrivateCards(userid)
	if err != nil {
		return nil, M.NewErrMsg(M.ErrCard.NotFound, err)
	}
	return data, nil
}

// GetCard 根据 ID 获取卡片
func GetCard(c *gin.Context) (models.Card, error) {
	id, err := strconv.Atoi(c.Param("cardid"))
	if err != nil {
		return models.Card{}, M.ErrCard.IDIllegal
	}
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return models.Card{}, M.ErrUser.IDIllegal
	}
	data, err := models.GetPrivateCardByID(userid, id)
	if err != nil {
		return models.Card{}, M.ErrCard.NotFound
	}
	return data, nil
}

// ModCard 修改卡片
func ModCard(c *gin.Context) error {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return M.ErrUser.IDIllegal
	}
	id, err := strconv.Atoi(c.Param("cardid"))
	if err != nil {
		return M.ErrCard.IDIllegal
	}

	question := c.PostForm("question")
	answer := c.PostForm("answer")
	if _, err = models.GetPrivateCardByID(userid, id); err != nil {
		return M.ErrCard.NotFound
	}
	if err = models.ModCardByID(id, question, answer); err != nil {
		return M.NewErrMsg(M.ErrCard.ModFail, err)
	}
	return nil
}

// DelCard 删除卡片
func DelCard(c *gin.Context) error {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return M.ErrUser.IDIllegal
	}
	id, err := strconv.Atoi(c.Param("cardid"))
	if err != nil {
		return M.ErrCard.IDIllegal
	}
	if _, err = models.GetPrivateCardByID(userid, id); err != nil {
		return M.NewErrMsg(M.ErrCard.NotFound, err)
	}
	if err = models.DelCardByID(id); err != nil {
		return M.NewErrMsg(M.ErrCard.DelFail, err)
	}
	return nil
}
