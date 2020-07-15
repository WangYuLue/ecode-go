package user

import (
	"net/http"
	"strconv"

	"ecode/models"
	"ecode/utils/message"

	"github.com/gin-gonic/gin"
)

// AddCard 添加卡片
func AddCard(c *gin.Context) {
	var u models.CardORM
	if c.ShouldBind(&u) != nil {
		message.HandelError(c, message.ErrHTTPData.BindFail)
		return
	}
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return
	}
	u.AutherID = userid
	if models.AddCard(&u) != nil {
		message.HandelError(c, message.ErrCard.ADDFail)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "添加卡片成功",
	})
}

// GetCards 获取所有卡片
func GetCards(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return
	}
	data, err := models.GetPrivateCards(userid)
	if err != nil {
		message.HandelError(c, message.ErrCard.NotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// GetCard 根据 ID 获取卡片
func GetCard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("cardid"))
	if err != nil {
		message.HandelError(c, message.ErrCard.IDIllegal)
		return
	}
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return
	}
	data, err := models.GetPrivateCardByID(userid, id)
	if err != nil {
		message.HandelError(c, message.ErrCard.NotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// ModCard 修改卡片
func ModCard(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return
	}
	id, err := strconv.Atoi(c.Param("cardid"))
	if err != nil {
		message.HandelError(c, message.ErrCard.IDIllegal)
		return
	}
	question := c.PostForm("question")
	answer := c.PostForm("answer")

	_, err = models.GetPrivateCardByID(userid, id)
	if err != nil {
		message.HandelError(c, message.ErrCard.NotFound)
		return
	}
	err = models.ModCardByID(id, question, answer)
	if err != nil {
		message.HandelError(c, message.ErrCard.ModFail)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "修改成功",
	})
}

// DelCard 删除卡片
func DelCard(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return
	}
	id, err := strconv.Atoi(c.Param("cardid"))
	if err != nil {
		message.HandelError(c, message.ErrCard.IDIllegal)
		return
	}
	_, err = models.GetPrivateCardByID(userid, id)
	if err != nil {
		message.HandelError(c, message.ErrCard.NotFound)
		return
	}
	err = models.DelCardByID(id)
	if err != nil {
		message.HandelError(c, message.ErrCard.DelFail)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}
