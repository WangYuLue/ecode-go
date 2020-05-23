package databases

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//数据库配置
const (
	userName = "root"
	password = "123qwe"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "ecode"
)

// SQLDB db
var SQLDB *gorm.DB

func init() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8&parseTime=True"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	db, err := gorm.Open("mysql", path)
	if err != nil {
		panic("failed to connect database")
	}
	SQLDB = db
	fmt.Println("connnect success")
}
