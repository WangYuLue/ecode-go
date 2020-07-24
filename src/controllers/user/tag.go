package user

import (
	"net/http"
	"strconv"

	"ecode/models"
	M "ecode/utils/message"

	"github.com/gin-gonic/gin"
)

// AddTag 添加标签
func AddTag(c *gin.Context) error {
	var u models.TagORM
	if c.ShouldBind(&u) != nil {
		return M.ErrHTTPData.BindFail
	}
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return M.ErrUser.IDIllegal
	}
	u.AutherID = userid
	if models.AddTag(&u) != nil {
		return M.NewErrMsg(M.ErrTag.AddFail, err)
	}
	return nil
}

// GetTags 获取所有标签
func GetTags(c *gin.Context) ([]models.Tag, error) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return nil, M.ErrUser.IDIllegal
	}
	data, err := models.GetPrivateTags(userid)
	if err != nil {
		return nil, M.NewErrMsg(M.ErrTag.NotFound, err)
	}
	return data, nil
}

// GetTag 根据 ID 获取标签
func GetTag(c *gin.Context) (models.Tag, error) {
	id, err := strconv.Atoi(c.Param("tagid"))
	if err != nil {
		return models.Tag{}, M.ErrTag.IDIllegal
	}
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return models.Tag{}, M.ErrUser.IDIllegal
	}
	data, err := models.GetPrivateTagByID(userid, id)
	if err != nil {
		return models.Tag{}, M.NewErrMsg(M.ErrTag.NotFound, err)
	}
	return data, nil
}

// ModTag 修改标签
func ModTag(c *gin.Context) error {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return M.ErrUser.IDIllegal
	}
	id, err := strconv.Atoi(c.Param("tagid"))
	if err != nil {
		return M.ErrTag.IDIllegal
	}
	name := c.PostForm("name")

	if _, err = models.GetPrivateTagByID(userid, id); err != nil {
		return M.NewErrMsg(M.ErrTag.NotFound, err)
	}
	if err = models.ModTagByID(id, name); err != nil {
		return M.NewErrMsg(M.ErrTag.ModFail, err)
	}
	return nil
}

// DelTag 删除标签
func DelTag(c *gin.Context) error {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return M.ErrUser.IDIllegal
	}
	id, err := strconv.Atoi(c.Param("tagid"))
	if err != nil {
		return M.ErrTag.IDIllegal
	}
	if _, err = models.GetPrivateTagByID(userid, id); err != nil {
		return M.NewErrMsg(M.ErrTag.NotFound, err)
	}
	if err = models.DelTagByID(id); err != nil {
		return M.NewErrMsg(M.ErrTag.DelFail, err)
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
