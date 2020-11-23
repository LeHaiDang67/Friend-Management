package controller

import (
	"encoding/json"
	"fmt"
	"friend_management/internal/db"
	"friend_management/internal/feature"
	"friend_management/internal/feature/model"
	"friend_management/internal/feature/user"
	"friend_management/pkg/ihttp"
	"net/http"
)

//GetUser is...
func GetUser(db db.Executor) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		if email == "" {
			ihttp.Respond(w, http.StatusBadRequest, feature.ResponseError{
				Error:       "missing_input_param",
				Description: "Error missing email param",
			})
			return
		}

		user, err := user.GetUser(db, email)
		if err != nil {
			ihttp.Respond(w, http.StatusInternalServerError, feature.ResponseError{
				Error:       "error_external_server",
				Description: "Cannot fetch user",
			})
			return
		}
		ihttp.Respond(w, http.StatusOK, user)
	})
}

//GetAllUsers is...
func GetAllUsers(db db.Executor) http.HandlerFunc {
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

//ConnectFriends is...
func ConnectFriends(db db.Executor) http.HandlerFunc {
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
func FriendList(db db.Executor) http.HandlerFunc {
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
func CommonFriends(db db.Executor) http.HandlerFunc {
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
func Subscription(db db.Executor) http.HandlerFunc {
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
func Blocked(db db.Executor) http.HandlerFunc {
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
func SendUpdate(db db.Executor) http.HandlerFunc {
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
