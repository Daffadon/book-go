package main

import (
	"book/db"
	"book/server"
)

func main() {
	db.Init()
	server.Init()
}
