package user

import (
	"net/http"
	"strconv"

	"ecode/models"
	"ecode/utils/message"

	"github.com/gin-gonic/gin"
)

// AddCategory 添加分类
func AddCategory(c *gin.Context) {
	var u models.CategoryORM
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
	if models.AddCategory(&u) != nil {
		message.HandelError(c, message.ErrCategory.ADDFail)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "发布分类成功",
	})
}

// GetCategorys 获取所有分类
func GetCategorys(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return
	}
	data, err := models.GetPrivateCategorys(userid)
	if err != nil {
		message.HandelError(c, message.ErrCategory.NotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// GetCategory 根据 ID 获取分类
func GetCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("categoryid"))
	if err != nil {
		message.HandelError(c, message.ErrCategory.IDIllegal)
		return
	}
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return
	}
	data, err := models.GetPrivateCategoryByID(userid, id)
	if err != nil {
		message.HandelError(c, message.ErrCategory.NotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// ModCategory 修改分类
func ModCategory(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return
	}
	id, err := strconv.Atoi(c.Param("categoryid"))
	if err != nil {
		message.HandelError(c, message.ErrCategory.IDIllegal)
		return
	}
	name := c.PostForm("name")

	_, err = models.GetPrivateCategoryByID(userid, id)
	if err != nil {
		message.HandelError(c, message.ErrCategory.NotFound)
		return
	}
	err = models.ModCategoryByID(id, name)
	if err != nil {
		message.HandelError(c, message.ErrCategory.ModFail)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "修改成功",
	})
}

// DelCategory 删除分类
func DelCategory(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return
	}
	id, err := strconv.Atoi(c.Param("categoryid"))
	if err != nil {
		message.HandelError(c, message.ErrCategory.IDIllegal)
		return
	}
	_, err = models.GetPrivateCategoryByID(userid, id)
	if err != nil {
		message.HandelError(c, message.ErrCategory.NotFound)
		return
	}
	err = models.DelCategoryByID(id)
	if err != nil {
		message.HandelError(c, message.ErrCategory.DelFail)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

// AddCardToCategory -
func AddCardToCategory(c *gin.Context) {
	// cardid := c.PostForm("cardid")
	// categoryid := c.PostForm("categoryid")
	// TODO:
	c.JSON(http.StatusOK, gin.H{
		"message": "添加成功",
	})
}

// RemoveCardToCategory -
func RemoveCardToCategory(c *gin.Context) {
	// TODO:
	c.JSON(http.StatusOK, gin.H{
		"message": "移除成功",
	})
}
