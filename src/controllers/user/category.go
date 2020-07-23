package user

import (
	"fmt"
	"strconv"

	"ecode/models"
	M "ecode/utils/message"

	"github.com/gin-gonic/gin"
)

// AddCategory 添加分类
func AddCategory(c *gin.Context) error {
	var u models.CategoryORM
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
	if err = models.AddCategory(&u); err != nil {
		M.HandelError(c, M.NewErrMsg(M.ErrCategory.AddFail, err))
		return ErrorDefault
	}
	return nil
}

// GetCategorys 获取所有分类
func GetCategorys(c *gin.Context) ([]models.Category, error) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		M.HandelError(c, M.ErrUser.IDIllegal)
		return nil, err
	}
	data, err := models.GetPrivateCategorys(userid)
	if err != nil {
		M.HandelError(c, M.NewErrMsg(M.ErrCategory.NotFound, err))
		return nil, err
	}
	return data, nil
}

// GetCategory 根据 ID 获取分类
func GetCategory(c *gin.Context) (models.Category, error) {
	id, err := strconv.Atoi(c.Param("categoryid"))
	if err != nil {
		M.HandelError(c, M.ErrCategory.IDIllegal)
		return CategoryDefault, err
	}
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		M.HandelError(c, M.ErrUser.IDIllegal)
		return CategoryDefault, err
	}
	data, err := models.GetPrivateCategoryByID(userid, id)
	if err != nil {
		M.HandelError(c, M.NewErrMsg(M.ErrCategory.NotFound, err))
		return CategoryDefault, err
	}
	return data, nil
}

// ModCategory 修改分类
func ModCategory(c *gin.Context) error {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		M.HandelError(c, M.ErrUser.IDIllegal)
		return ErrorDefault
	}
	id, err := strconv.Atoi(c.Param("categoryid"))
	if err != nil {
		M.HandelError(c, M.ErrCategory.IDIllegal)
		return ErrorDefault
	}
	name := c.PostForm("name")

	_, err = models.GetPrivateCategoryByID(userid, id)
	if err != nil {
		M.HandelError(c, M.NewErrMsg(M.ErrCategory.NotFound, err))
		return ErrorDefault
	}
	err = models.ModCategoryByID(id, name)
	if err != nil {
		M.HandelError(c, M.NewErrMsg(M.ErrCategory.ModFail, err))
		return ErrorDefault
	}
	return nil
}

// DelCategory 删除分类
func DelCategory(c *gin.Context) error {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		M.HandelError(c, M.ErrUser.IDIllegal)
		return ErrorDefault
	}
	id, err := strconv.Atoi(c.Param("categoryid"))
	if err != nil {
		M.HandelError(c, M.ErrCategory.IDIllegal)
		return ErrorDefault
	}
	_, err = models.GetPrivateCategoryByID(userid, id)
	if err != nil {
		M.HandelError(c, M.NewErrMsg(M.ErrCategory.NotFound, err))
		return ErrorDefault
	}
	err = models.DelCategoryByID(id)
	if err != nil {
		M.HandelError(c, M.NewErrMsg(M.ErrCategory.DelFail, err))
		return ErrorDefault
	}
	return nil
}

// AddCardToCategory -
func AddCardToCategory(c *gin.Context) error {
	cardid, err := strconv.Atoi(c.PostForm("cardid"))
	categoryid, err := strconv.Atoi(c.PostForm("categoryid"))
	userid, err := strconv.Atoi(c.Param("userid"))
	// 验证参数是否合法
	if err != nil {
		M.HandelError(c, M.ErrHTTPData.Illegal)
		return ErrorDefault
	}
	// TODO: 验证是否是当前用户的 card 、category
	card, err := models.GetPrivateCardByID(userid, cardid)
	fmt.Println(card)
	fmt.Println(err)
	// 验证是否已经添加
	num := models.IsCardCategoryExist(cardid, categoryid)
	if num > 0 {
		M.HandelError(c, M.ErrCardCategory.HasAdd)
		return ErrorDefault
	}
	err = models.AddCardToCategory(cardid, categoryid)
	if err != nil {
		M.HandelError(c, M.NewErrMsg(M.ErrHTTPData.AddFail, err))
		return ErrorDefault
	}
	return nil
}

// RemoveCardToCategory -
func RemoveCardToCategory(c *gin.Context) error {
	cardid, err := strconv.Atoi(c.PostForm("cardid"))
	categoryid, err := strconv.Atoi(c.PostForm("categoryid"))
	if err != nil {
		M.HandelError(c, M.ErrHTTPData.Illegal)
		return ErrorDefault
	}
	err = models.RemoveCardToCategory(cardid, categoryid)
	if err != nil {
		M.HandelError(c, M.NewErrMsg(M.ErrHTTPData.DelFail, err))
		return ErrorDefault
	}
	return nil
}
