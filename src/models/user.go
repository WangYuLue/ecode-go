package models

import (
	"time"
)

// User 用户对象
type User struct {
	ID           int       `gorm:"column:user_id;AUTO_INCREMENT;primary_key" json:"id" form:"id"`
	Name         string    `gorm:"type:varchar(100);NOT NULL;DEFAULT:''" json:"name" form:"name"`
	Introduction string    `gorm:"type:varchar(1000);NOT NULL;DEFAULT:''"`
	Github       string    `gorm:"type:varchar(100);NOT NULL;DEFAULT:''"`
	Password     string    `gorm:"type:varchar(100);NOT NULL;DEFAULT:''"`
	PersonURL    string    `gorm:"type:varchar(100);NOT NULL;DEFAULT:''"`
	Email        string    `gorm:"type:varchar(100);NOT NULL;DEFAULT:''"`
	CreatedAt    time.Time `gorm:"NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	// 用户状态 -1 表示已删除； 0 表示未激活； 1 表示正常
	Status int    `gorm:"NOT NULL;DEFAULT:0"`
	Data   string `gorm:"DEFAULT:NULL"`
	Cards  []Card `gorm:"foreignkey:AutherID"`
}

// Login 用户登录
func Login(email, password string) (user User, err error) {
	queryString := "status = ? AND email = ? AND password = ? AND email <> '' AND password <> ''"
	err = SQLDB.Where(queryString, 0, email, password).Find(&user).Error
	return
}

// TableName 表名
func (User) TableName() string {
	return "user"
}

// AddUser 添加 user
func AddUser(p *User) (user User, err error) {
	err = SQLDB.Create(p).Find(&user).Error
	return
}

// GetUsers 获取所有 user
func GetUsers() (users []User, err error) {
	err = SQLDB.Not(User{Status: -1}).Find(&users).Error
	return
}

// GetUserByID 根据 ID 获取 user
func GetUserByID(id int) (user User, err error) {
	err = SQLDB.Not(User{Status: -1}).Find(&user, id).Error
	// 关联查询：
	// 第一种方法用 Related 有效,
	// 第二种方法用 Related 有效, 而且这种方式不用写 `gorm:"foreignkey:AutherID"` 就能查询
	// 第三种方法用 Association 有效,
	// 第四种方法用 Preload 有效, 可以一次性查询
	// SQLDB.Debug().Model(&user).Related(&user.Cards, "Cards")
	// SQLDB.Debug().Model(&user).Related(&user.Cards, "AutherID")
	// SQLDB.Debug().Model(&user).Association("Cards").Find(&user.Cards)
	// SQLDB.Debug().Not(User{Status: -1}).Preload("Cards").Find(&user, id).Error
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
func ModUserByID(id int, name string) (err error) {
	err = SQLDB.Model(User{ID: id}).Update("name", name).Error
	return
}

// DelUserByID 根据 ID 删除 user
func DelUserByID(id int) (err error) {
	err = SQLDB.Model(User{ID: id}).Update("status", -1).Error
	return
}
