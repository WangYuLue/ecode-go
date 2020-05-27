package models

import (
	mysql "ecode/databases/mysql"
)

// SQLDB orm 层
var SQLDB = mysql.DB

// Mail 邮件格式
type Mail struct {
	Name string
	URL  string
}
