package user

import (
	"fmt"
	"friend_management/internal/db"
	"friend_management/internal/feature/model"
	"friend_management/internal/feature/util"
	"log"
	"strings"

	"github.com/lib/pq"
)

//ConnectFriends that func connect 2 user
func ConnectFriends(db db.Executor, req model.FriendConnectionRequest) (model.BasicResponse, error) {
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

	bBlock := util.Contains(userA.Blocked, userB.Email)
	aBlock := util.Contains(userB.Blocked, userA.Email)
	if aBlock || bBlock {
		basicResponse.Success = false
		return basicResponse, nil
	}

	bFriend := util.Contains(userA.Friends, userB.Email)
	aFriend := util.Contains(userB.Friends, userA.Email)
	if !bFriend || !aFriend {
		errUpdateA := AddFriends(db, userB.Email, userA.Email)
		if errUpdateA != nil {
			fmt.Printf("Error QueryA: %s\n", errUpdateA)
		}
		log.Printf("B added to A friend's\n")
		errUpdateB := AddFriends(db, userA.Email, userB.Email)
		if errUpdateB != nil {
			fmt.Printf("Error QueryB: %s\n", errUpdateB)
		}
		log.Printf("A added to B friend's\n")
	}

	basicResponse.Success = true
	return basicResponse, nil
}

//FriendList show friend list
func FriendList(db db.Executor, email string) (model.FriendListResponse, error) {
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
func CommonFriends(db db.Executor, commonFriends model.CommonFriendRequest) (model.FriendListResponse, error) {
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

//Subscription subscribe to updates from an email address.
func Subscription(db db.Executor, subRequest model.SubscriptionRequest) (model.BasicResponse, error) {
	var basicResponse model.BasicResponse
	userRequestor, errGetUser1 := GetUser(db, subRequest.Requestor)
	if errGetUser1 != nil {
		basicResponse.Success = false
		return basicResponse, errGetUser1
	}
	userTarget, errGetUser2 := GetUser(db, subRequest.Target)
	if errGetUser2 != nil {
		basicResponse.Success = false
		return basicResponse, errGetUser2
	}
	isUserRequestor := util.Contains(userRequestor.Subscription, userTarget.Email)
	if !isUserRequestor {
		result, err := db.Exec("Update users set subscription = array_append(subscription,$1)  where email = $2 ",
			userTarget.Email, userRequestor.Email)
		if err != nil {
			basicResponse.Success = false
			return basicResponse, err
		}
		result.RowsAffected()
	}

	basicResponse.Success = true
	return basicResponse, nil
}

//Blocked is  an API to block updates from an email address
func Blocked(db db.Executor, subRequest model.SubscriptionRequest) (model.BasicResponse, error) {
	var basicResponse model.BasicResponse
	userRequestor, errGetUser1 := GetUser(db, subRequest.Requestor)
	if errGetUser1 != nil {
		basicResponse.Success = false
		return basicResponse, errGetUser1
	}
	userTarget, errGetUser2 := GetUser(db, subRequest.Target)
	if errGetUser2 != nil {
		basicResponse.Success = false
		return basicResponse, errGetUser2
	}
	isUserRequestor := util.Contains(userRequestor.Blocked, userTarget.Email)
	if !isUserRequestor {
		result, errQuery := db.Exec("Update users set blocked = array_append(blocked,$1)  where email = $2 ",
			userTarget.Email, userRequestor.Email)
		if errQuery != nil {
			basicResponse.Success = false
			return basicResponse, errQuery
		}

		result.RowsAffected()
	}

	basicResponse.Success = true
	return basicResponse, nil

}

//SendUpdate retrieve all email addresses that can receive updates from an email address.
func SendUpdate(db db.Executor, sendRequest model.SendUpdateRequest) (model.SendUpdateResponse, error) {
	var sendResponse model.SendUpdateResponse
	sender, err1 := GetUser(db, sendRequest.Sender)
	if err1 != nil {
		sendResponse.Success = false
		return sendResponse, nil
	}
	Recipients := []string{}
	allUser, err2 := GetAllUsers(db)
	if err2 != nil {
		sendResponse.Success = false
		return sendResponse, nil
	}
	for _, u := range allUser {
		var isBlock = util.Contains(u.Blocked, sender.Email)
		if !isBlock {
			isFriend := util.Contains(u.Friends, sender.Email)
			isSubscriber := util.Contains(u.Subscription, sender.Email)
			isMentioned := strings.Contains(sendRequest.Text, u.Email)
			if isFriend || isSubscriber || isMentioned {
				Recipients = append(Recipients, u.Email)
			}

		}
	}
	sendResponse.Success = true
	sendResponse.Recipients = Recipients
	return sendResponse, nil
}

//GetUser get user bu email
func GetUser(db db.Executor, email string) (model.User, error) {
	user := model.User{}
	query := "select * from users where email = $1"
	r, err := db.Query(query, email)
	if err != nil {
		return user, err
	}

	for r.Next() {
		err := r.Scan(&user.Email, pq.Array(&user.Friends), pq.Array(&user.Subscription), pq.Array(&user.Blocked))
		if err != nil {
			return user, err
		}
	}

	return user, nil
}

//GetAllUsers get all user
func GetAllUsers(db db.Executor) ([]model.User, error) {
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

//AddFriends add a new friend
func AddFriends(db db.Executor, emailFriend string, email string) error {
	result, err := db.Exec("Update users set friends=array_append(friends,$1)  where email = $2 ",
		emailFriend, email)
	if err != nil {
		return err
	}

	result.RowsAffected()
	return nil
}
