package controller

import (
	"database/sql"
	"encoding/json"
	"friend_management/intenal/feature/repo"
	"net/http"
)

//GetUser is...
func GetUser(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")

		user, err := repo.GetUser(db, email)
		if err != nil {
			json.NewEncoder(w).Encode("Cannot fetch user")
			return
		}
		json.NewEncoder(w).Encode(user)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("HTTP status code returned!"))
	})
}
