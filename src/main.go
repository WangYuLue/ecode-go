package main

import (
	db "ecode/databases"
	"ecode/routers"
)

func main() {

	defer db.SQLDB.Close()
	router := routers.InitRouter()
	router.Run(":8000")
}
