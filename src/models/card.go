package models

import (
	"time"
)

// Card 卡片对象
type Card struct {
	ID        int       `gorm:"column:card_id;AUTO_INCREMENT;primary_key" json:"id" form:"id"`
	Question  string    `gorm:"type:varchar(100);NOT NULL" json:"question" form:"question"`
	Answer    string    `gorm:"type:text;NOT NULL" json:"answer" form:"answer"`
	AutherID  int       `gorm:"type:varchar(1000);NOT NULL" json:"auther_id" form:"auther_id"`
	CreatedAt time.Time `gorm:"NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	IsDeleted int       `gorm:"NOT NULL;DEFAULT:0"`
	Data      string    `gorm:"DEFAULT:NULL"`
}

// TableName 表名
func (Card) TableName() string {
	return "card"
}

// AddCard 添加 card
func AddCard(p *Card) (err error) {
	err = SQLDB.Create(p).Error
	return
}

// GetCards 获取所有 card
func GetCards() (cards []Card, err error) {
	err = SQLDB.Where("status = ?", 0).Find(&cards).Error
	return
}

// GetCardByID 根据 ID 修改 card
func GetCardByID(id int) (card Card, err error) {
	err = SQLDB.Where("status = ?", 0).Find(&card, id).Error
	return
}

// ModCardByID 根据 ID 修改 card
func ModCardByID(id int, question string, answer string) (err error) {
	err = SQLDB.Model(Card{ID: id}).Update(Card{Question: question, Answer: answer}).Error
	return
}

// DelCardByID 根据 ID 删除 card
func DelCardByID(id int) (err error) {
	err = SQLDB.Model(Card{ID: id}).Update("status", -1).Error
	return
}
