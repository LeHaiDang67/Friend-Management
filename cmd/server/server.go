package server

import (
	"fmt"
	"friend_management/cmd/router"
	"friend_management/internal/db"
	"net/http"
	"os"
)

// Start starts the application server
func Start(db db.Executor) {
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
