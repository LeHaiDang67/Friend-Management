package server

import (
	"database/sql"
	"fmt"
	"friend_management/cmd/router"
	"net/http"
	"os"
)

// Start starts the application server
func Start(db *sql.DB) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router.Handler(db),
	}
	server.ListenAndServe()
}
