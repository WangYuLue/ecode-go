package mysql

import (
	"ecode/config"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// DB db
var DB *gorm.DB

func init() {
	mysql := config.Mysql
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	// loc=Local 设置正确时区
	path := strings.Join([]string{mysql.UserName, ":", mysql.Password, "@tcp(", mysql.IP, ":", mysql.Port, ")/", mysql.DBName, "?charset=utf8&parseTime=True&loc=Local"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	db, err := gorm.Open("mysql", path)
	if err != nil {
		panic("mysql 连接失败")
	}
	DB = db
	fmt.Println("mysql 连接成功")
}
