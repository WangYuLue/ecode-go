package v1

import (
	"net/http"
	"strconv"

	"ecode/models"
	"ecode/utils/message"

	"github.com/gin-gonic/gin"
)

// AddCardAPI 添加卡片
func AddCardAPI(c *gin.Context) {
	var u models.CardORM
	if c.ShouldBind(&u) != nil {
		message.HandelError(c, message.ErrHTTPData.BindFail)
		return
	}
	if models.AddCard(&u) != nil {
		message.HandelError(c, message.ErrCard.ADDFail)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "添加卡片成功",
	})
}

// GetCardsAPI 获取所有卡片
func GetCardsAPI(c *gin.Context) {
	data, err := models.GetCards()
	if err != nil {
		message.HandelError(c, message.ErrCard.NotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// GetCardAPI 根据 ID 获取卡片
func GetCardAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("cardid"))
	if err != nil {
		message.HandelError(c, message.ErrCard.IDIllegal)
		return
	}
	data, err := models.GetCardByID(id)
	if err != nil {
		message.HandelError(c, message.ErrCard.NotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// ModCardAPI 修改卡片
func ModCardAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("cardid"))
	question := c.PostForm("question")
	answer := c.PostForm("answer")
	if err != nil {
		message.HandelError(c, message.ErrCard.IDIllegal)
		return
	}
	_, err = models.GetCardByID(id)
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

// DelCardAPI 删除卡片
func DelCardAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("cardid"))
	if err != nil {
		message.HandelError(c, message.ErrCard.IDIllegal)
		return
	}
	_, err = models.GetCardByID(id)
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
