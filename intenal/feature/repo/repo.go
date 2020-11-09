package repo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"friend_management/intenal/feature/model"
	"friend_management/intenal/feature/util"
	"log"

	"github.com/lib/pq"
)

//FakeUser is a simgle user
type FakeUser struct {
	Email        string `json:"email"`
	Friends      string `json:"friends"`
	Subscription string `json:"subscription"`
	Blocked      string `json:"blocked"`
}

//ConnectFriends that func connect 2 user
func ConnectFriends(db *sql.DB, req model.FriendConnectionRequest) (model.BasicResponse, error) {
	basicResponse := model.BasicResponse{}

	userA, errA := GetUser(db, req.Friends[0])
	userB, errB := GetUser(db, req.Friends[1])
	if errA != nil {
		fmt.Printf("Error QueryA: %s\n", errA)
		basicResponse.Success = false
		return basicResponse, errA
	}
	if errB != nil {
		fmt.Printf("Error QueryB: %s\n", errB)
		basicResponse.Success = false
		return basicResponse, errB
	}
	singleUserA := changeSingleUser(userA)
	singleUserB := changeSingleUser(userB)

	bBlock := util.Contains(userA.Blocked, userB.Email)
	aBlock := util.Contains(userB.Blocked, userA.Email)
	if aBlock || bBlock {
		basicResponse.Success = false
		return basicResponse, nil
	}

	bFriend := util.Contains(userA.Friends, userB.Email)
	aFriend := util.Contains(userB.Friends, userA.Email)
	if !bFriend || !aFriend {
		errUpdateA := UpdateUser2(db, singleUserB, userA.Email)
		if errUpdateA != nil {
			fmt.Printf("Error QueryA: %s\n", errUpdateA)
		}
		log.Printf("B added to A friend's\n")
		errUpdateB := UpdateUser2(db, singleUserA, userB.Email)
		if errUpdateB != nil {
			fmt.Printf("Error QueryB: %s\n", errUpdateB)
		}
		log.Printf("A added to B friend's\n")
	}

	basicResponse.Success = true
	return basicResponse, nil
}

//FriendList show friend list
func FriendList(db *sql.DB, email string) (model.FriendListResponse, error) {
	user := model.User{}
	var friendList model.FriendListResponse
	r, err1 := db.Query("select * from users where email = $1", email)
	if err1 != nil {
		friendList.Success = false
		return friendList, err1
	}
	for r.Next() {
		err := r.Scan(&user.Email, pq.Array(&user.Friends), pq.Array(&user.Subscription), pq.Array(&user.Blocked))
		if err != nil {
			friendList.Success = false
			return friendList, err
		}
	}
	friendList.Success = true
	friendList.Friends = user.Friends
	friendList.Count = len(user.Friends)
	return friendList, nil
}

//CommonFriends retrieve the common friends list between two email addresses
func CommonFriends(db *sql.DB, commonFriends model.CommonFriendRequest) (model.FriendListResponse, error) {
	var friendList model.FriendListResponse
	userA, errA := GetUser(db, commonFriends.Friends[0])
	userB, errB := GetUser(db, commonFriends.Friends[1])
	if errA != nil {
		fmt.Printf("Error QueryA: %s\n", errA)
		friendList.Success = false
		return friendList, errA
	}
	if errB != nil {
		fmt.Printf("Error QueryB: %s\n", errB)
		friendList.Success = false
		return friendList, errB
	}
	Commons := []string{}
	for _, a := range userA.Friends {
		for _, b := range userB.Friends {
			if a == b {
				Commons = append(Commons, a)
			}
		}
	}
	friendList.Success = true
	friendList.Friends = Commons
	friendList.Count = len(Commons)
	return friendList, nil
}

//GetUser get user bu email
func GetUser(db *sql.DB, email string) (model.User, error) {
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

//UpdateUser edit the user
func UpdateUser(db *sql.DB, user FakeUser, email string) error {

	result, err := db.Exec("Update users set friends=array[$1] , subscription = array[$2], blocked = array[$3] where email = $4 ",
		user.Friends, user.Subscription, user.Blocked, email)
	if err != nil {
		return err
	}

	result.RowsAffected()
	return nil
}

//UpdateUser2 append the user []
func UpdateUser2(db *sql.DB, user FakeUser, email string) error {

	result, err := db.Exec("Update users set friends=array_append(friends,$1)  where email = $2 ",
		email, user.Email)
	if err != nil {
		return err
	}

	result.RowsAffected()
	return nil
}

func changeSingleUser(arrUser model.User) FakeUser {
	uFriend, errFriend := json.Marshal(arrUser.Friends)
	uFriendStr := string(uFriend)
	if errFriend != nil {
		panic(errFriend)
	}
	uSub, errSub := json.Marshal(arrUser.Subscription)
	uSubStr := string(uSub)
	if errSub != nil {
		panic(errSub)
	}
	uBlocked, errBlocked := json.Marshal(arrUser.Blocked)
	uBlockedStr := string(uBlocked)
	if errBlocked != nil {
		panic(errFriend)
	}
	uFake := FakeUser{Email: arrUser.Email,
		Friends:      uFriendStr,
		Subscription: uSubStr,
		Blocked:      uBlockedStr}

	return uFake
}
