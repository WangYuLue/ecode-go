package user

import (
	"strconv"

	"ecode/models"
	"ecode/utils/message"

	"github.com/gin-gonic/gin"
)

// AddCategory 添加分类
func AddCategory(c *gin.Context) error {
	var u models.CategoryORM
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
	if models.AddCategory(&u) != nil {
		message.HandelError(c, message.ErrCategory.AddFail)
		return ErrorDefault
	}
	return nil
}

// GetCategorys 获取所有分类
func GetCategorys(c *gin.Context) ([]models.Category, error) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return nil, err
	}
	data, err := models.GetPrivateCategorys(userid)
	if err != nil {
		message.HandelError(c, message.ErrCategory.NotFound)
		return nil, err
	}
	return data, nil
}

// GetCategory 根据 ID 获取分类
func GetCategory(c *gin.Context) (models.Category, error) {
	id, err := strconv.Atoi(c.Param("categoryid"))
	if err != nil {
		message.HandelError(c, message.ErrCategory.IDIllegal)
		return CategoryDefault, err
	}
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return CategoryDefault, err
	}
	data, err := models.GetPrivateCategoryByID(userid, id)
	if err != nil {
		message.HandelError(c, message.ErrCategory.NotFound)
		return CategoryDefault, err
	}
	return data, nil
}

// ModCategory 修改分类
func ModCategory(c *gin.Context) error {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return ErrorDefault
	}
	id, err := strconv.Atoi(c.Param("categoryid"))
	if err != nil {
		message.HandelError(c, message.ErrCategory.IDIllegal)
		return ErrorDefault
	}
	name := c.PostForm("name")

	_, err = models.GetPrivateCategoryByID(userid, id)
	if err != nil {
		message.HandelError(c, message.ErrCategory.NotFound)
		return ErrorDefault
	}
	err = models.ModCategoryByID(id, name)
	if err != nil {
		message.HandelError(c, message.ErrCategory.ModFail)
		return ErrorDefault
	}
	return nil
}

// DelCategory 删除分类
func DelCategory(c *gin.Context) error {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		message.HandelError(c, message.ErrUser.IDIllegal)
		return ErrorDefault
	}
	id, err := strconv.Atoi(c.Param("categoryid"))
	if err != nil {
		message.HandelError(c, message.ErrCategory.IDIllegal)
		return ErrorDefault
	}
	_, err = models.GetPrivateCategoryByID(userid, id)
	if err != nil {
		message.HandelError(c, message.ErrCategory.NotFound)
		return ErrorDefault
	}
	err = models.DelCategoryByID(id)
	if err != nil {
		message.HandelError(c, message.ErrCategory.DelFail)
		return ErrorDefault
	}
	return nil
}

// AddCardToCategory -
func AddCardToCategory(c *gin.Context) error {
	cardid, err := strconv.Atoi(c.PostForm("cardid"))
	categoryid, err := strconv.Atoi(c.PostForm("categoryid"))
	// 验证参数是否合法
	if err != nil {
		message.HandelError(c, message.ErrHTTPData.Illegal)
		return ErrorDefault
	}
	// TODO: 验证是否是当前用户的 card 、category

	// 验证是否已经添加
	num := models.IsCardCategoryExist(cardid, categoryid)
	if num > 0 {
		message.HandelError(c, message.ErrCardCategory.HasAdd)
		return ErrorDefault
	}
	err = models.AddCardToCategory(cardid, categoryid)
	if err != nil {
		message.HandelError(c, message.ErrHTTPData.AddFail)
		return ErrorDefault
	}
	return nil
}

// RemoveCardToCategory -
func RemoveCardToCategory(c *gin.Context) error {
	cardid, err := strconv.Atoi(c.PostForm("cardid"))
	categoryid, err := strconv.Atoi(c.PostForm("categoryid"))
	if err != nil {
		message.HandelError(c, message.ErrHTTPData.Illegal)
		return ErrorDefault
	}
	err = models.RemoveCardToCategory(cardid, categoryid)
	if err != nil {
		message.HandelError(c, message.ErrHTTPData.DelFail)
		return ErrorDefault
	}
	return nil
}
