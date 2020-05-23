package models

import (
	"time"
)

// Card 卡片对象
type Card struct {
	ID        int       `gorm:"column:card_id;AUTO_INCREMENT;primary_key" json:"id" form:"id"`
	Question  string    `gorm:"type:varchar(100);NOT NULL" json:"question" form:"question"`
	Answer    string    `gorm:"type:text;NOT NULL" json:"answer" form:"answer"`
	AutherID  int       `gorm:"type:varchar(1000);NOT NULL"`
	CreatedAt time.Time `gorm:"NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	IsDeleted int       `gorm:"NOT NULL;DEFAULT:0"`
	Data      string    `gorm:"DEFAULT:NULL"`
}

// TableName 表名
func (Card) TableName() string {
	return "card"
}

// AddCard 添加卡片
func AddCard(p *Card) (id int64, err error) {
	err = SQLDB.Create(&Card{Question: p.Question}).Error
	return
}

// GetCards 获取所有卡片
func GetCards() (cards []Card, err error) {
	err = SQLDB.Where("is_deleted = ?", 0).Find(&cards).Error
	return
}

// GetCardByID 根据 ID 修改卡片
func GetCardByID(id int) (card Card, err error) {
	err = SQLDB.Where("is_deleted = ?", 0).Find(&card, id).Error
	return
}

// ModCardByID 根据 ID 修改卡片
func ModCardByID(id int, question string, answer string) (err error) {
	err = SQLDB.Model(Card{ID: id}).Update(Card{Question: question, Answer: answer}).Error
	return
}

// DelCardByID 根据 ID 删除卡片
func DelCardByID(id int) (err error) {
	err = SQLDB.Model(Card{ID: id}).Update("is_deleted", -1).Error
	return
}
