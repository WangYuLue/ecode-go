package user

import (
	"net/http"
	"strconv"

	"ecode/models"
	"ecode/utils/message"

	"github.com/gin-gonic/gin"
)

// AddTag 添加标签
func AddTag(c *gin.Context) error {
	var u models.TagORM
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
	if models.AddTag(&u) != nil {
		message.HandelError(c, message.ErrTag.AddFail)
		return ErrorDefault
	}
	return nil
}

// GetTags 获取所有标签
func GetTags(c *gin.Context) ([]models.Tag, error) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return nil, err
	}
	data, err := models.GetPrivateTags(userid)
	if err != nil {
		message.HandelError(c, message.ErrTag.NotFound)
		return nil, err
	}
	return data, nil
}

// GetTag 根据 ID 获取标签
func GetTag(c *gin.Context) (models.Tag, error) {
	id, err := strconv.Atoi(c.Param("tagid"))
	if err != nil {
		message.HandelError(c, message.ErrTag.IDIllegal)
		return TagDefault, err
	}
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return TagDefault, err
	}
	data, err := models.GetPrivateTagByID(userid, id)
	if err != nil {
		message.HandelError(c, message.ErrTag.NotFound)
		return TagDefault, err
	}
	return data, nil
}

// ModTag 修改标签
func ModTag(c *gin.Context) error {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return ErrorDefault
	}
	id, err := strconv.Atoi(c.Param("tagid"))
	if err != nil {
		message.HandelError(c, message.ErrTag.IDIllegal)
		return ErrorDefault
	}
	name := c.PostForm("name")

	_, err = models.GetPrivateTagByID(userid, id)
	if err != nil {
		message.HandelError(c, message.ErrTag.NotFound)
		return ErrorDefault
	}
	err = models.ModTagByID(id, name)
	if err != nil {
		message.HandelError(c, message.ErrTag.ModFail)
		return ErrorDefault
	}
	return nil
}

// DelTag 删除标签
func DelTag(c *gin.Context) error {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return ErrorDefault
	}
	id, err := strconv.Atoi(c.Param("tagid"))
	if err != nil {
		message.HandelError(c, message.ErrTag.IDIllegal)
		return ErrorDefault
	}
	_, err = models.GetPrivateTagByID(userid, id)
	if err != nil {
		message.HandelError(c, message.ErrTag.NotFound)
		return ErrorDefault
	}
	err = models.DelTagByID(id)
	if err != nil {
		message.HandelError(c, message.ErrTag.DelFail)
		return ErrorDefault
	}
	return nil
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
