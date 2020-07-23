package models

import (
	"errors"
	"time"
)

const categotyTableName = "category"

// CategoryORM 映射
type CategoryORM struct {
	ID        int       `gorm:"column:category_id;AUTO_INCREMENT;primary_key" json:"id" form:"id"`
	Name      string    `gorm:"type:varchar(100);NOT NULL" json:"name" form:"name"`
	AutherID  int       `gorm:"type:int;NOT NULL" json:"auther_id" form:"auther_id"`
	CreatedAt time.Time `gorm:"NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	Status    int       `gorm:"NOT NULL;DEFAULT:0"`
	Data      string    `gorm:"DEFAULT:NULL"`
}

// Category -
type Category struct {
	ID        int       `gorm:"column:category_id" json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	AutherID  int       `json:"auther_id" form:"auther_id"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt"`
}

// TableName 表名
func (CategoryORM) TableName() string {
	return categotyTableName
}

// AddCategory 添加 category
func AddCategory(p *CategoryORM) (err error) {
	err = SQLDB.Create(p).Error
	return
}

// GetPublicCategoryByID 根据 ID 获取【公开的】 category
func GetPublicCategoryByID(id int) (category Category, err error) {
	err = SQLDB.Table(categotyTableName).Where("status = ?", "0").Where(CategoryORM{ID: id}).Scan(&category).Error
	if err == nil && category.ID == 0 {
		err = errors.New("catagory not find")
	}
	return
}

// GetPrivateCategorys 获取所有【私有的】 category
func GetPrivateCategorys(autherID int) (category []Category, err error) {
	err = SQLDB.Table(categotyTableName).Where(CategoryORM{AutherID: autherID}).Where("status <> ?", "-1").Scan(&category).Error
	return
}

// GetPrivateCategoryByID 根据 ID 获取【私有的】 category
func GetPrivateCategoryByID(autherID int, id int) (category Category, err error) {
	err = SQLDB.Table(categotyTableName).Where(CategoryORM{AutherID: autherID}).Where("status = ?", "0").Where(CategoryORM{ID: id}).Scan(&category).Error
	if err == nil && category.ID == 0 {
		err = errors.New("catagory not find")
	}
	return
}

// ModCategoryByID 根据 ID 修改 category
func ModCategoryByID(id int, name string) (err error) {
	err = SQLDB.Model(CategoryORM{ID: id}).Update(CategoryORM{Name: name}).Error
	return
}

// DelCategoryByID 根据 ID 删除 category
func DelCategoryByID(id int) (err error) {
	err = SQLDB.Model(CategoryORM{ID: id}).Update("status", -1).Error
	return
}
