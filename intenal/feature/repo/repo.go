package repo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"friend_management/intenal/feature/model"

	"github.com/lib/pq"
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

	// sqlStatement := "select * from users where email = $1"

	// stmt, err := db.Prepare(sqlStatement)
	// if err != nil {
	// 	// handle err
	// }

	for r.Next() {
		err := r.Scan(&user.Email, pq.Array(user.Friends), pq.Array(&user.Subscription), pq.Array(&user.Blocked))
		if err != nil {
			panic(err)
		}
	}
	us, errs := json.Marshal(user)
	if errs != nil {

	}
	fmt.Println(string(us))
	return user, nil
}
