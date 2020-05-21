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
	CreatedAt    time.Time `gorm:"NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	IsDeleted    int       `gorm:"NOT NULL;DEFAULT:0"`
	Data         string    `gorm:"DEFAULT:NULL"`
}

// TableName 表名
func (User) TableName() string {
	return "user"
}

// AddUser 添加用户
func AddUser(p *User) (id int64, err error) {
	err = SQLDB.Create(&User{Nick: p.Nick}).Error
	return
}

// GetUsers 获取所有用户
func GetUsers() (users []User, err error) {
	err = SQLDB.Where(&User{IsDeleted: 0}).Find(&users).Error
	return
}

// GetUserByID 根据 ID 获取用户
func GetUserByID(id int) (user User, err error) {
	err = SQLDB.Where(&User{IsDeleted: 0}).Find(&user, id).Error
	return
}

// ModUserByID 根据 ID 修改用户
func ModUserByID(id int, nick string) (err error) {
	err = SQLDB.Model(User{ID: id}).Update("nick", nick).Error
	return
}

// DelUserByID 根据 ID 删除用户
func DelUserByID(id int) (err error) {
	err = SQLDB.Model(User{ID: id}).Update("is_deleted", -1).Error
	return
}
