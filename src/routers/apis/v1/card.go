package v1

import (
	"net/http"
	"strconv"

	"ecode/models"
	"ecode/utils"

	"github.com/gin-gonic/gin"
)

// AddCardAPI 添加卡片
func AddCardAPI(c *gin.Context) {
	var u models.Card
	if c.ShouldBind(&u) != nil {
		utils.HandelError(c, utils.StatusBadMessage.Illegal.Data)
		return
	}
	if models.AddCard(&u) != nil {
		utils.HandelError(c, utils.StatusBadMessage.Fail.Add)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "添加卡片成功",
	})
}

// GetCardsAPI 获取所有卡片
func GetCardsAPI(c *gin.Context) {
	data, err := models.GetCards()
	if err != nil {
		utils.HandelError(c, utils.StatusBadMessage.Fail.Get)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// GetCardAPI 根据 ID 获取卡片
func GetCardAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.HandelError(c, utils.StatusBadMessage.Illegal.ID)
		return
	}
	data, err := models.GetCardByID(id)
	if err != nil {
		utils.HandelError(c, utils.StatusBadMessage.None.Card)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// ModCardAPI 修改卡片
func ModCardAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	question := c.PostForm("question")
	answer := c.PostForm("answer")
	if err != nil {
		utils.HandelError(c, utils.StatusBadMessage.Illegal.ID)
		return
	}
	_, err = models.GetCardByID(id)
	if err != nil {
		utils.HandelError(c, utils.StatusBadMessage.None.Card)
		return
	}
	err = models.ModCardByID(id, question, answer)
	if err != nil {
		utils.HandelError(c, utils.StatusBadMessage.Fail.Mod)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "修改成功",
	})
}

// DelCardAPI 删除卡片
func DelCardAPI(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.HandelError(c, utils.StatusBadMessage.Illegal.ID)
		return
	}
	_, err = models.GetCardByID(id)
	if err != nil {
		utils.HandelError(c, utils.StatusBadMessage.None.Card)
		return
	}
	err = models.DelCardByID(id)
	if err != nil {
		utils.HandelError(c, utils.StatusBadMessage.Fail.Del)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "删除成功",
	})
}
