package models

import (
	mysql "ecode/databases/mysql"
)

// SQLDB orm 层
var SQLDB = mysql.DB

// Login -
func Login(email, password string) (user User, err error) {
	queryString := "status = ? AND email = ? AND password = ? AND email <> '' AND password <> ''"
	err = SQLDB.Where(queryString, 0, email, password).Find(&user).Error
	return
}
