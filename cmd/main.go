package main

import (
	"friend_management/cmd/server"
	"friend_management/intenal/db"
)

func main() {
	db := db.InitDatabase()
	defer db.Close()

	// Start server
	server.Start(db)
}
