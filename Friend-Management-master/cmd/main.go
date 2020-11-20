package main

import (
	"friend_management/cmd/server"
	"friend_management/internal/db"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := db.InitDatabase()
	defer db.Close()

	// Start server
	server.Start(db)
}
