package main

import (
	"github.com/inoue0124/salon-be/db"
	"github.com/inoue0124/salon-be/server"
)

func main() {
	db.Init()
	server.Init()
}
