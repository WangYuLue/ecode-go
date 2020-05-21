package main

import (
	db "ecode/databases"
)

func main() {
	defer db.SQLDB.Close()
	router := initRouter()
	router.Run(":8000")
}
