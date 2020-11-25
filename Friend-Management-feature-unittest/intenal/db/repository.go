package db

import (
	"database/sql"
	"friend_management/intenal/model"

	"github.com/lib/pq"
)

//GetUserByEmail executes retrieve the friends list for an email address
func GetUserByEmail(db *sql.DB, email string) (*model.User, error) {
	user := &model.User{}
	var friendList model.FriendListResponse
	r, err1 := db.Query("select * from users where email = $1", email)
	if err1 != nil {
		friendList.Success = false
		return user, err1
	}
	for r.Next() {
		err := r.Scan(&user.Email, pq.Array(&user.Friends), pq.Array(&user.Subscription), pq.Array(&user.Blocked))
		if err != nil {
			friendList.Success = false
			return user, err
		}
	}
	return user, nil

}

//GetTheUser is used to get api
func GetTheUser(db *sql.DB, email string) (model.User, error) {
	user := model.User{}

	r, err1 := db.Query("select * from users where email = $1", email)
	if err1 != nil {
		return user, err1
	}

	for r.Next() {
		err := r.Scan(&user.Email, pq.Array(&user.Friends), pq.Array(&user.Subscription), pq.Array(&user.Blocked))
		if err != nil {
			return user, err
		}
	}

	return user, nil
}

//GetListUsers is used to get API
func GetListUsers(db *sql.DB) ([]model.User, error) {
	users := []model.User{}
	user := model.User{}
	r, err1 := db.Query("select * from users")
	if err1 != nil {
		return users, err1
	}
	for r.Next() {
		err := r.Scan(&user.Email, pq.Array(&user.Friends), pq.Array(&user.Subscription), pq.Array(&user.Blocked))
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

// CreateSubscribeFriend executes subscribe to updates from an email address
func CreateSubscribeFriend(db *sql.DB, userRequestor, userTarget string) error {
	result, err := db.Exec("Update users set subscription = array_append(subscription,$1)  where email = $2 ",
		userTarget, userRequestor)
	if err != nil {
		return err
	}
	result.RowsAffected()
	return nil
}

//CreateBlockedFriend is ...
func CreateBlockedFriend(db *sql.DB, userRequestor, userTarget string) error {
	result, errQuery := db.Exec("Update users set blocked = array_append(blocked,$1)  where email = $2 ",
		userTarget, userRequestor)
	if errQuery != nil {
		return errQuery
	}

	result.RowsAffected()
	return nil
}
