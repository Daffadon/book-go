package main

import (
	"bookApp/db"
	"bookApp/server"
)

func main() {
	db.Init()
	server.Init()
}
