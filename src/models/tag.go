package models

import (
	"errors"
	"time"
)

const tagTableName = "tag"

// TagORM 映射
type TagORM struct {
	ID        int       `gorm:"column:tag_id;AUTO_INCREMENT;primary_key" json:"id" form:"id"`
	Name      string    `gorm:"type:varchar(100);NOT NULL" json:"name" form:"name"`
	AutherID  int       `gorm:"type:int;NOT NULL" json:"auther_id" form:"auther_id"`
	CreatedAt time.Time `gorm:"NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	Status    int       `gorm:"NOT NULL;DEFAULT:0"`
	Data      string    `gorm:"DEFAULT:NULL"`
}

// Tag -
type Tag struct {
	ID        int       `gorm:"column:tag_id" json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	AutherID  int       `json:"auther_id" form:"auther_id"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt"`
}

// TableName 表名
func (TagORM) TableName() string {
	return tagTableName
}

// AddTag 添加 tag
func AddTag(p *TagORM) (err error) {
	err = SQLDB.Create(p).Error
	return
}

// GetPublicTagByID 根据 ID 获取【公开的】 tag
func GetPublicTagByID(id int) (tag Tag, err error) {
	err = SQLDB.Table(tagTableName).Where("status = ?", "0").Where(TagORM{ID: id}).Scan(&tag).Error
	if err == nil && tag.ID == 0 {
		err = errors.New("tag not find")
	}
	return
}

// GetPrivateTags 获取所有【私有的】 tag
func GetPrivateTags(autherID int) (tag []Tag, err error) {
	err = SQLDB.Table(tagTableName).Where(TagORM{AutherID: autherID}).Where("status <> ?", "-1").Scan(&tag).Error
	return
}

// GetPrivateTagByID 根据 ID 获取【私有的】 tag
func GetPrivateTagByID(autherID int, id int) (tag Tag, err error) {
	err = SQLDB.Table(tagTableName).Where(TagORM{AutherID: autherID}).Where("status = ?", "0").Where(TagORM{ID: id}).Scan(&tag).Error
	if err == nil && tag.ID == 0 {
		err = errors.New("tag not find")
	}
	return
}

// ModTagByID 根据 ID 修改 tag
func ModTagByID(id int, name string) (err error) {
	err = SQLDB.Model(TagORM{ID: id}).Update(TagORM{Name: name}).Error
	return
}

// DelTagByID 根据 ID 删除 tag
func DelTagByID(id int) (err error) {
	err = SQLDB.Model(TagORM{ID: id}).Update("status", -1).Error
	return
}
