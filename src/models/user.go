package models

import (
	"time"
)

// User 用户对象
type User struct {
	ID           int       `gorm:"column:user_id;AUTO_INCREMENT;primary_key" json:"id" form:"id"`
	Nick         string    `gorm:"type:varchar(100);NOT NULL;DEFAULT:''" json:"nick" form:"nick"`
	Introduction string    `gorm:"type:varchar(1000);NOT NULL;DEFAULT:''"`
	Github       string    `gorm:"type:varchar(100);NOT NULL;DEFAULT:''"`
	PersonURL    string    `gorm:"type:varchar(100);NOT NULL;DEFAULT:''"`
	Email        string    `gorm:"type:varchar(100);NOT NULL;DEFAULT:''"`
	CreatedAt    time.Time `gorm:"NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	IsDeleted    int       `gorm:"NOT NULL;DEFAULT:0"`
	Data         string    `gorm:"DEFAULT:NULL"`
	Cards        []Card    `gorm:"foreignkey:AutherID"`
}

// TableName 表名
func (User) TableName() string {
	return "user"
}

// AddUser 添加 user
func AddUser(p *User) (err error) {
	err = SQLDB.Create(p).Error
	return
}

// GetUsers 获取所有 user
func GetUsers() (users []User, err error) {
	err = SQLDB.Not(User{IsDeleted: -1}).Find(&users).Error
	return
}

// GetUserByID 根据 ID 获取 user
func GetUserByID(id int) (user User, err error) {
	err = SQLDB.Not(User{IsDeleted: -1}).Find(&user, id).Error
	// 关联查询：
	// 第一种方法用 Related 有效,
	// 第二种方法用 Related 有效, 而且这种方式不用写 `gorm:"foreignkey:AutherID"` 就能查询
	// 第三种方法用 Association 有效,
	// 第四种方法用 Preload 有效, 可以一次性查询
	// SQLDB.Debug().Model(&user).Related(&user.Cards, "Cards")
	// SQLDB.Debug().Model(&user).Related(&user.Cards, "AutherID")
	// SQLDB.Debug().Model(&user).Association("Cards").Find(&user.Cards)
	// SQLDB.Debug().Not(User{IsDeleted: -1}).Preload("Cards").Find(&user, id).Error
	// 参考文章：https://segmentfault.com/a/1190000017263285
	return
}

// GetCardsByUserID 根据 ID 获取 user 下的 card
func GetCardsByUserID(id int) (cards []Card, err error) {
	_, err = GetUserByID(id)
	if err != nil {
		return
	}
	err = SQLDB.Where(Card{AutherID: id}).Find(&cards).Error
	if err != nil {
		return
	}
	return
}

// ModUserByID 根据 ID 修改 user
func ModUserByID(id int, nick string) (err error) {
	err = SQLDB.Model(User{ID: id}).Update("nick", nick).Error
	return
}

// DelUserByID 根据 ID 删除 user
func DelUserByID(id int) (err error) {
	err = SQLDB.Model(User{ID: id}).Update("is_deleted", -1).Error
	return
}
