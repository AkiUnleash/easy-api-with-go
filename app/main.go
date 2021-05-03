package main

import (
	"crudapi/server"

	_ "crudapi/docs"

	_ "github.com/go-sql-driver/mysql"
)

// @title Easy API with Go
// @version 1.0.0
// @description This is program I developed as a lerning process before creating an API in Golang.

// @license.name MIT
// @license.url https://github.com/tcnksm/tool/blob/master/LICENCE

// @host http://localhost:8081
// @BasePath /

func main() {

	db, err := server.SqlConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	server.Routing(db)

}
