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
		M.HandelError(c, M.ErrHTTPData.BindFail)
		return ErrorDefault
	}
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		M.HandelError(c, M.ErrUser.IDIllegal)
		return ErrorDefault
	}
	u.AutherID = userid
	if err = models.AddCard(&u); err != nil {
		M.HandelError(c, M.NewErrMsg(M.ErrCard.AddFail, err))
		return ErrorDefault
	}
	return nil
}

// GetCards 获取所有卡片
func GetCards(c *gin.Context) ([]models.Card, error) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		M.HandelError(c, M.ErrUser.IDIllegal)
		return nil, err
	}
	data, err := models.GetPrivateCards(userid)
	if err != nil {
		M.HandelError(c, M.NewErrMsg(M.ErrCard.NotFound, err))
		return nil, err
	}
	return data, nil
}

// GetCard 根据 ID 获取卡片
func GetCard(c *gin.Context) (models.Card, error) {
	id, err := strconv.Atoi(c.Param("cardid"))
	if err != nil {
		M.HandelError(c, M.ErrCard.IDIllegal)
		return CardDefault, err
	}
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		M.HandelError(c, M.ErrUser.IDIllegal)
		return CardDefault, err
	}
	// TODO: 将 err 提取到外面
	data, err := models.GetPrivateCardByID(userid, id)
	if err != nil {
		M.HandelError(c, M.ErrCard.NotFound)
		return CardDefault, err
	}
	return data, nil
}

// ModCard 修改卡片
func ModCard(c *gin.Context) error {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		M.HandelError(c, M.ErrUser.IDIllegal)
		return ErrorDefault
	}
	id, err := strconv.Atoi(c.Param("cardid"))
	if err != nil {
		M.HandelError(c, M.ErrCard.IDIllegal)
		return ErrorDefault
	}
	question := c.PostForm("question")
	answer := c.PostForm("answer")

	_, err = models.GetPrivateCardByID(userid, id)
	if err != nil {
		M.HandelError(c, M.ErrCard.NotFound)
		return ErrorDefault
	}
	err = models.ModCardByID(id, question, answer)
	if err != nil {
		M.HandelError(c, M.NewErrMsg(M.ErrCard.ModFail, err))
		return ErrorDefault
	}
	return nil
}

// DelCard 删除卡片
func DelCard(c *gin.Context) error {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		M.HandelError(c, M.ErrUser.IDIllegal)
		return ErrorDefault
	}
	id, err := strconv.Atoi(c.Param("cardid"))
	if err != nil {
		M.HandelError(c, M.ErrCard.IDIllegal)
		return ErrorDefault
	}
	_, err = models.GetPrivateCardByID(userid, id)
	if err != nil {
		M.HandelError(c, M.NewErrMsg(M.ErrCard.NotFound, err))
		return ErrorDefault
	}
	err = models.DelCardByID(id)
	if err != nil {
		M.HandelError(c, M.NewErrMsg(M.ErrCard.DelFail, err))
		return ErrorDefault
	}
	return nil
}
