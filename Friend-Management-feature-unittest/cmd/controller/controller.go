package controller

import (
	"encoding/json"
	"fmt"

	"friend_management/intenal/model"
	"friend_management/intenal/services"

	"net/http"
)

//GetUser is...
func GetUser(service services.IUserService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")

		user, err := service.GetUser(email)
		if err != nil {
			json.NewEncoder(w).Encode("Cannot fetch user")
			return
		}
		json.NewEncoder(w).Encode(user)
		w.WriteHeader(http.StatusOK)

	})
}

//GetAllUsers is...
func GetAllUsers(service services.IUserService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		listUser, err := service.GetAllUsers()
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
func ConnectFriends(service services.IUserService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req model.FriendConnectionRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Request body is invalid"))
			return
		}
		basicResponse, err1 := service.ConnectFriends(&req)
		if err1 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err1.Error()))
			return
		}
		json.NewEncoder(w).Encode(basicResponse)
	})
}

//FriendList is...
func FriendList(service services.IUserService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		friendList, err := service.FriendList(email)
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
func CommonFriends(service services.IUserService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var commonFriends model.CommonFriendRequest
		err := json.NewDecoder(r.Body).Decode(&commonFriends)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Request body is invalid"))
			return
		}
		friendList, err1 := service.CommonFriends(commonFriends)
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
func Subscription(service services.IUserService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var subRequest model.SubscriptionRequest
		err := json.NewDecoder(r.Body).Decode(&subRequest)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Request body is invalid"))
			return
		}
		result, errSub := service.Subscription(subRequest)
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
func Blocked(service services.IUserService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var subRequest model.SubscriptionRequest
		err := json.NewDecoder(r.Body).Decode(&subRequest)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Request body is invalid"))
			return
		}
		result, errSub := service.Blocked(subRequest)
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
func SendUpdate(service services.IUserService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var sendRequest model.SendUpdateRequest
		err := json.NewDecoder(r.Body).Decode(&sendRequest)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Request body is invalid"))
			return
		}
		result, err2 := service.SendUpdate(sendRequest)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err2.Error()))
			return
		}
		json.NewEncoder(w).Encode(result)
		w.WriteHeader(http.StatusOK)

	})
}
