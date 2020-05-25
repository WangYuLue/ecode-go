package models

import (
	db "ecode/databases"
)

// SQLDB orm 层
var SQLDB = db.SQLDB

// Login -
func Login(email, password string) (user User, err error) {
	queryString := "is_deleted = ? AND email = ? AND password = ? AND email <> '' AND password <> ''"
	err = SQLDB.Where(queryString, 0, email, password).Find(&user).Error
	return
}
