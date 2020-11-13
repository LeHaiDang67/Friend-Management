package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"friend_management/intenal/feature/model"
	"friend_management/intenal/feature/user"
	"net/http"
)

//GetUser is...
func GetUser(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")

		user, err := user.GetUser(db, email)
		if err != nil {
			json.NewEncoder(w).Encode("Cannot fetch user")
			return
		}
		json.NewEncoder(w).Encode(user)
		w.WriteHeader(http.StatusOK)

	})
}

//GetAllUsers is...
func GetAllUsers(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		listUser, err := user.GetAllUsers(db)
		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode("Cannot fetch users")
			return
		}
		json.NewEncoder(w).Encode(listUser)
		w.WriteHeader(http.StatusOK)

	})
}

<<<<<<< HEAD
=======
//UpdateUser is...
func UpdateUser(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		var requestUser model.User

		err := json.NewDecoder(r.Body).Decode(&requestUser)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Request body is invalid"))
			return
		}

		errs := user.UpdateUser(db, requestUser, email)
		if errs != nil {

			json.NewEncoder(w).Encode("Cannot update user")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("HTTP status code returned!"))
	})
}

>>>>>>> b65464a... Moving user_test to internal
//ConnectFriends is...
func ConnectFriends(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var u model.User

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Request body is invalid"))
			return
		}
		req := model.FriendConnectionRequest{Friends: u.Friends}
		basicResponse, err1 := user.ConnectFriends(db, req)
		if err1 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err1.Error()))
			return
		}
		json.NewEncoder(w).Encode(basicResponse)
	})
}

//FriendList is...
func FriendList(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		friendList, err := user.FriendList(db, email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		json.NewEncoder(w).Encode(friendList)
		w.WriteHeader(http.StatusOK)

	})
}

//CommonFriends is...
func CommonFriends(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var commonFriends model.CommonFriendRequest
		err := json.NewDecoder(r.Body).Decode(&commonFriends)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Request body is invalid"))
			return
		}
		friendList, err1 := user.CommonFriends(db, commonFriends)
		if err1 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err1.Error()))
			return
		}
		json.NewEncoder(w).Encode(friendList)
		w.WriteHeader(http.StatusOK)

	})
}

//Subscription is...
func Subscription(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var subRequest model.SubscriptionRequest
		err := json.NewDecoder(r.Body).Decode(&subRequest)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Request body is invalid"))
			return
		}
		result, errSub := user.Subscription(db, subRequest)
		if errSub != nil {
			json.NewEncoder(w).Encode(errSub)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errSub.Error()))
			return
		}
		json.NewEncoder(w).Encode(result)
		w.WriteHeader(http.StatusOK)

	})
}

//Blocked is...
func Blocked(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var subRequest model.SubscriptionRequest
		err := json.NewDecoder(r.Body).Decode(&subRequest)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Request body is invalid"))
			return
		}
		result, errSub := user.Blocked(db, subRequest)
		if errSub != nil {
			json.NewEncoder(w).Encode(errSub)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errSub.Error()))
			return
		}
		json.NewEncoder(w).Encode(result)
		w.WriteHeader(http.StatusOK)

	})
}

// SendUpdate is ...
func SendUpdate(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var sendRequest model.SendUpdateRequest
		err := json.NewDecoder(r.Body).Decode(&sendRequest)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Request body is invalid"))
			return
		}
		result, err2 := user.SendUpdate(db, sendRequest)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err2.Error()))
			return
		}
		json.NewEncoder(w).Encode(result)
		w.WriteHeader(http.StatusOK)

	})
}
