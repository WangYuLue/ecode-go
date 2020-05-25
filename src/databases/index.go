package databases

import (
	"ecode/config"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// SQLDB db
var SQLDB *gorm.DB

func init() {
	mysql := config.Mysql
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{mysql.UserName, ":", mysql.Password, "@tcp(", mysql.IP, ":", mysql.Port, ")/", mysql.DBName, "?charset=utf8&parseTime=True"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	db, err := gorm.Open("mysql", path)
	if err != nil {
		panic("failed to connect database")
	}
	SQLDB = db
	fmt.Println("connnect success")
}
