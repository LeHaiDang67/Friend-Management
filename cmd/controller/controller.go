package controller

import (
	"database/sql"
	"encoding/json"
	"friend_management/intenal/feature/model"
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

//UpdateUser is...
func UpdateUser(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		var requestUser repo.FakeUser

		err := json.NewDecoder(r.Body).Decode(&requestUser)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("HTTP status code returned!"))
			return
		}

		errs := repo.UpdateUser(db, requestUser, email)
		if errs != nil {

			json.NewEncoder(w).Encode("Cannot update user")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("HTTP status code returned!"))
	})
}

//ConnectFriends is...
func ConnectFriends(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var u model.User

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("HTTP status code returned!"))
			return
		}
		req := model.FriendConnectionRequest{Friends: u.Friends}
		basicResponse, err1 := repo.ConnectFriends(db, req)
		if err1 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("HTTP status code returned!"))
			return
		}
		json.NewEncoder(w).Encode(basicResponse)
	})
}

//FriendList is...
func FriendList(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		friendList, err := repo.FriendList(db, email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("HTTP status code returned!"))
			return
		}
		json.NewEncoder(w).Encode(friendList)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("HTTP status code returned!"))
	})
}

//CommonFriends is...
func CommonFriends(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var commonFriends model.CommonFriendRequest
		err := json.NewDecoder(r.Body).Decode(&commonFriends)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("HTTP status code returned!"))
			return
		}
		friendList, err1 := repo.CommonFriends(db, commonFriends)
		if err1 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("HTTP status code returned!"))
			return
		}
		json.NewEncoder(w).Encode(friendList)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("HTTP status code returned!"))
	})
}

//Subscription is...
func Subscription(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var subRequest model.SubscriptionRequest
		err := json.NewDecoder(r.Body).Decode(&subRequest)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("HTTP status code returned!"))
			return
		}
		result := repo.Subscription(db, subRequest)
		json.NewEncoder(w).Encode(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("HTTP status code returned!"))
	})
}
