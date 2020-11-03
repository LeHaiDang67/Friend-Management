package repo

import (
	"database/sql"
	"friend_management/intenal/feature/model"
)

//ConnectFriends that func connect 2 user
func ConnectFriends(db *sql.DB, email1, email2 string) model.BasicResponse {
	basicResponse := model.BasicResponse{}

	basicResponse.Success = true
	return basicResponse
}

//GetUser get user bu email
func GetUser(db *sql.DB, email string) (model.User, error) {
	user := model.User{}
	r, err := db.Query("select * from users where email = $1", email)
	if err != nil {
		return user, err
	}
	//json.Marshal(r)
	for r.Next() {
		err := r.Scan(&user.Email, &user.Friends, &user.Subscription, &user.Blocked)
		if err != nil {
			panic(err)
		}
	}
	return user, nil
}
