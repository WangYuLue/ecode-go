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
		return M.ErrHTTPData.BindFail
	}
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return M.ErrUser.IDIllegal
	}
	u.AutherID = userid
	if err = models.AddCategory(&u); err != nil {
		return M.NewErrMsg(M.ErrCategory.AddFail, err)
	}
	return nil
}

// GetCategorys 获取所有分类
func GetCategorys(c *gin.Context) ([]models.Category, error) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return nil, M.ErrUser.IDIllegal
	}
	data, err := models.GetPrivateCategorys(userid)
	if err != nil {
		return nil, M.NewErrMsg(M.ErrCategory.NotFound, err)
	}
	return data, nil
}

// GetCategory 根据 ID 获取分类
func GetCategory(c *gin.Context) (models.Category, error) {
	id, err := strconv.Atoi(c.Param("categoryid"))
	if err != nil {
		return models.Category{}, M.ErrCategory.IDIllegal
	}
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return models.Category{}, M.ErrUser.IDIllegal
	}
	data, err := models.GetPrivateCategoryByID(userid, id)
	if err != nil {
		return models.Category{}, M.NewErrMsg(M.ErrCategory.NotFound, err)
	}
	return data, nil
}

// ModCategory 修改分类
func ModCategory(c *gin.Context) error {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return M.ErrUser.IDIllegal
	}
	id, err := strconv.Atoi(c.Param("categoryid"))
	if err != nil {
		return M.ErrCategory.IDIllegal
	}
	name := c.PostForm("name")
	if _, err = models.GetPrivateCategoryByID(userid, id); err != nil {
		return M.NewErrMsg(M.ErrCategory.NotFound, err)
	}
	if err = models.ModCategoryByID(id, name); err != nil {
		return M.NewErrMsg(M.ErrCategory.ModFail, err)
	}
	return nil
}

// DelCategory 删除分类
func DelCategory(c *gin.Context) error {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return M.ErrUser.IDIllegal
	}
	id, err := strconv.Atoi(c.Param("categoryid"))
	if err != nil {
		return M.ErrCategory.IDIllegal
	}
	if _, err = models.GetPrivateCategoryByID(userid, id); err != nil {
		return M.NewErrMsg(M.ErrCategory.NotFound, err)
	}
	if err = models.DelCategoryByID(id); err != nil {
		return M.NewErrMsg(M.ErrCategory.DelFail, err)
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
		return M.ErrHTTPData.Illegal
	}
	// TODO: 验证是否是当前用户的 card 、category
	card, err := models.GetPrivateCardByID(userid, cardid)
	fmt.Println(card)
	fmt.Println(err)
	// 验证是否已经添加
	if num := models.IsCardCategoryExist(cardid, categoryid); num > 0 {
		return M.ErrCardCategory.HasAdd
	}
	if err = models.AddCardToCategory(cardid, categoryid); err != nil {
		return M.NewErrMsg(M.ErrHTTPData.AddFail, err)
	}
	return nil
}

// RemoveCardToCategory -
func RemoveCardToCategory(c *gin.Context) error {
	cardid, err := strconv.Atoi(c.PostForm("cardid"))
	categoryid, err := strconv.Atoi(c.PostForm("categoryid"))
	if err != nil {
		return M.ErrHTTPData.Illegal
	}
	if err = models.RemoveCardToCategory(cardid, categoryid); err != nil {
		return M.NewErrMsg(M.ErrHTTPData.DelFail, err)
	}
	return nil
}
