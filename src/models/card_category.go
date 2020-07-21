package models

import (
	"time"
)

const cardCategotyTableName = "card_category"

// CardCategoryORM 映射
type CardCategoryORM struct {
	ID         int       `gorm:"AUTO_INCREMENT;primary_key" json:"id" form:"id"`
	CardID     int       `gorm:"type:int;NOT NULL" json:"card_id" form:"card_id"`
	CategoryID int       `gorm:"type:int;NOT NULL" json:"category_id" form:"category_id"`
	CreatedAt  time.Time `gorm:"NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	Data       string    `gorm:"DEFAULT:NULL"`
}

// CardCategory -
type CardCategory struct {
	ID         int       `json:"id" form:"id"`
	CardID     int       `json:"card_id" form:"card_id"`
	CategoryID int       `json:"category_id" form:"category_id"`
	CreatedAt  time.Time `json:"createdAt" form:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" form:"updatedAt"`
}

// TableName 表名
func (CardCategoryORM) TableName() string {
	return cardCategotyTableName
}

// AddCardToCategory -
func AddCardToCategory(cardID int, categoryID int) (err error) {
	err = SQLDB.Create(&CardCategoryORM{CardID: cardID, CategoryID: categoryID}).Error
	return
}

// RemoveCardToCategory -
func RemoveCardToCategory(cardID int, categoryID int) (err error) {
	err = SQLDB.Delete(CardCategoryORM{}, "card_id = ? AND category_id = ?", cardID, categoryID).Error
	return
}

// IsCardCategoryExist -
func IsCardCategoryExist(cardID int, categoryID int) (count int) {
	SQLDB.Model(CardCategoryORM{}).Where("card_id = ? AND category_id = ?", cardID, categoryID).Count(&count)
	return
}

// GetCardIDsByCategoryID -
func GetCardIDsByCategoryID(categoryID int) (cardids []int, err error) {
	// TODO:
	// err = SQLDB.Table(cardCategotyTableName).Where(CardCategoryORM{CategoryID: categoryID}).Scan(&cards).Error
	return
}

// GetCategoryIDsByCardID -
func GetCategoryIDsByCardID(categoryID int) (cards []Card, err error) {
	return
}
