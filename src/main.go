package main

import (
	"ecode/databases/mysql"
	"ecode/databases/redis"
	"ecode/routers"
)

func main() {
	defer mysql.DB.Close()
	defer redis.DB.Close()
	router := routers.InitRouter()
	router.Run(":8000")
}
