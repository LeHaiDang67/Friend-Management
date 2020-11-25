package services

import (
	"database/sql"
	"fmt"
	"friend_management/intenal/db"
	"friend_management/intenal/model"
	"friend_management/intenal/util"

	"log"
	"strings"
)

//IUserService is...
type IUserService interface {
	ConnectFriends(req *model.FriendConnectionRequest) (*model.BasicResponse, error)
	GetUser(email string) (model.User, error)
	SendUpdate(sendRequest model.SendUpdateRequest) (model.SendUpdateResponse, error)
	GetAllUsers() ([]model.User, error)
	Blocked(subRequest model.SubscriptionRequest) (model.BasicResponse, error)
	Subscription(subRequest model.SubscriptionRequest) (model.BasicResponse, error)
	CommonFriends(commonFriends model.CommonFriendRequest) (model.FriendListResponse, error)
	FriendList(email string) (model.FriendListResponse, error)
}

//Store is...
type Store struct {
	dbconn *sql.DB
}

//NewManager is...
func NewManager(dbconn *sql.DB) *Store {
	return &Store{
		dbconn: dbconn,
	}
}

//ConnectFriends that func connect 2 user
func (st *Store) ConnectFriends(req *model.FriendConnectionRequest) (*model.BasicResponse, error) {
	basicResponse := &model.BasicResponse{}
	//var service IUserService
	userA, errA := db.GetTheUser(st.dbconn, req.Friends[0])
	userB, errB := db.GetTheUser(st.dbconn, req.Friends[1])
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
		errUpdateA := AddFriends(st.dbconn, userB.Email, userA.Email)
		if errUpdateA != nil {
			fmt.Printf("Error QueryA: %s\n", errUpdateA)
		}
		log.Printf("B added to A friend's\n")
		errUpdateB := AddFriends(st.dbconn, userA.Email, userB.Email)
		if errUpdateB != nil {
			fmt.Printf("Error QueryB: %s\n", errUpdateB)
		}
		log.Printf("A added to B friend's\n")
	}

	basicResponse.Success = true
	return basicResponse, nil
}

//FriendList show friend list
func (st *Store) FriendList(email string) (model.FriendListResponse, error) {
	var friendList model.FriendListResponse
	user, err := db.GetUserByEmail(st.dbconn, email)
	if err != nil {
		return friendList, nil
	}
	friendList.Success = true
	friendList.Friends = user.Friends
	friendList.Count = len(user.Friends)
	return friendList, nil
}

//CommonFriends retrieve the common friends list between two email addresses
func (st *Store) CommonFriends(commonFriends model.CommonFriendRequest) (model.FriendListResponse, error) {
	var friendList model.FriendListResponse
	userA, errA := db.GetTheUser(st.dbconn, commonFriends.Friends[0])
	userB, errB := db.GetTheUser(st.dbconn, commonFriends.Friends[1])
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
func (st *Store) Subscription(subRequest model.SubscriptionRequest) (model.BasicResponse, error) {
	var basicResponse model.BasicResponse
	userRequestor, errGetUser1 := db.GetTheUser(st.dbconn, subRequest.Requestor)
	if errGetUser1 != nil {
		basicResponse.Success = false
		return basicResponse, errGetUser1
	}
	userTarget, errGetUser2 := db.GetTheUser(st.dbconn, subRequest.Target)
	if errGetUser2 != nil {
		basicResponse.Success = false
		return basicResponse, errGetUser2
	}
	isUserRequestor := util.Contains(userRequestor.Subscription, userTarget.Email)
	if !isUserRequestor {
		err := db.CreateSubscribeFriend(st.dbconn, userRequestor.Email, userTarget.Email)
		if err != nil {
			basicResponse.Success = false
			return basicResponse, err
		}
	}

	basicResponse.Success = true
	return basicResponse, nil
}

//Blocked is  an API to block updates from an email address
func (st *Store) Blocked(subRequest model.SubscriptionRequest) (model.BasicResponse, error) {
	var basicResponse model.BasicResponse
	userRequestor, errGetUser1 := db.GetTheUser(st.dbconn, subRequest.Requestor)
	if errGetUser1 != nil {
		basicResponse.Success = false
		return basicResponse, errGetUser1
	}
	userTarget, errGetUser2 := db.GetTheUser(st.dbconn, subRequest.Target)
	if errGetUser2 != nil {
		basicResponse.Success = false
		return basicResponse, errGetUser2
	}
	isUserRequestor := util.Contains(userRequestor.Blocked, userTarget.Email)
	if !isUserRequestor {
		err := db.CreateBlockedFriend(st.dbconn, userRequestor.Email, userTarget.Email)
		if err != nil {
			basicResponse.Success = false
			return basicResponse, err
		}
	}

	basicResponse.Success = true
	return basicResponse, nil

}

//SendUpdate retrieve all email addresses that can receive updates from an email address.
func (st *Store) SendUpdate(sendRequest model.SendUpdateRequest) (model.SendUpdateResponse, error) {
	var sendResponse model.SendUpdateResponse
	sender, err1 := db.GetTheUser(st.dbconn, sendRequest.Sender)
	if err1 != nil {
		sendResponse.Success = false
		return sendResponse, nil
	}
	Recipients := []string{}
	allUser, err2 := db.GetListUsers(st.dbconn)
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
func (st *Store) GetUser(email string) (model.User, error) {
	user, err := db.GetTheUser(st.dbconn, email)
	if err != nil {
		return user, err
	}
	return user, nil
}

//GetAllUsers get all user
func (st *Store) GetAllUsers() ([]model.User, error) {
	users, err := db.GetListUsers(st.dbconn)
	if err != nil {
		return users, err
	}
	return users, nil
}

//AddFriends add a new friend
func AddFriends(db *sql.DB, emailFriend string, email string) error {

	result, err := db.Exec("Update users set friends=array_append(friends,$1)  where email = $2 ",
		emailFriend, email)
	if err != nil {
		return err
	}

	result.RowsAffected()
	return nil
}
