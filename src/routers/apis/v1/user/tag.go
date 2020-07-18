package user

import (
	"net/http"
	"strconv"

	"ecode/models"
	"ecode/utils/message"

	"github.com/gin-gonic/gin"
)

// AddTag 添加标签
func AddTag(c *gin.Context) {
	var u models.TagORM
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
	if models.AddTag(&u) != nil {
		message.HandelError(c, message.ErrTag.ADDFail)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "发布标签成功",
	})
}

// GetTags 获取所有标签
func GetTags(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return
	}
	data, err := models.GetPrivateTags(userid)
	if err != nil {
		message.HandelError(c, message.ErrTag.NotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// GetTag 根据 ID 获取标签
func GetTag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("tagid"))
	if err != nil {
		message.HandelError(c, message.ErrTag.IDIllegal)
		return
	}
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return
	}
	data, err := models.GetPrivateTagByID(userid, id)
	if err != nil {
		message.HandelError(c, message.ErrTag.NotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// ModTag 修改标签
func ModTag(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return
	}
	id, err := strconv.Atoi(c.Param("tagid"))
	if err != nil {
		message.HandelError(c, message.ErrTag.IDIllegal)
		return
	}
	name := c.PostForm("name")

	_, err = models.GetPrivateTagByID(userid, id)
	if err != nil {
		message.HandelError(c, message.ErrTag.NotFound)
		return
	}
	err = models.ModTagByID(id, name)
	if err != nil {
		message.HandelError(c, message.ErrTag.ModFail)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "修改成功",
	})
}

// DelTag 删除标签
func DelTag(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return
	}
	id, err := strconv.Atoi(c.Param("tagid"))
	if err != nil {
		message.HandelError(c, message.ErrTag.IDIllegal)
		return
	}
	_, err = models.GetPrivateTagByID(userid, id)
	if err != nil {
		message.HandelError(c, message.ErrTag.NotFound)
		return
	}
	err = models.DelTagByID(id)
	if err != nil {
		message.HandelError(c, message.ErrTag.DelFail)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

// AddCardToTag -
func AddCardToTag(c *gin.Context) {
	// TODO:
	c.JSON(http.StatusOK, gin.H{
		"message": "添加成功",
	})
}

// RemoveCardToTag -
func RemoveCardToTag(c *gin.Context) {
	// TODO:
	c.JSON(http.StatusOK, gin.H{
		"message": "移除成功",
	})
}
