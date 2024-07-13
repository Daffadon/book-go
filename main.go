package main

import (
	"bookApp/db"
	"bookApp/migrations"
	"bookApp/server"
)

func main() {
	db.Init()
	migrations.Init()
	server.Init()
}
