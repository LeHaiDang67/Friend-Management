package user

import (
	"errors"
	"friend_management/internal/db"
	"friend_management/internal/feature"
	"friend_management/internal/feature/model"
	"friend_management/pkg/testutil"
	"testing"

	"github.com/stretchr/testify/require"
)

//TestGetUser test GetUser func
func TestGetUser(t *testing.T) {
	testCases := []struct {
		desc           string
		givenUserEmail string
		expectedResult model.User
		expectedError  error
	}{
		{
			desc:           "Should return user",
			givenUserEmail: "test-email@gmail.com",
			expectedResult: model.User{
				Email:   "test-email@gmail.com",
				Friends: []string{"hero@gmail.com"},
			},
			expectedError: nil,
		},
		{
			desc:           "Should return no user",
			givenUserEmail: "andy@example.com",
			expectedResult: model.User{},
			expectedError:  errors.New("sad"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.WithTxDB(t, func(tx db.BeginnerExecutor) {
				testutil.LoadTestDataFile(t, tx, "internal/feature/user/testData/01_get_user.sql")
				result, err := GetUser(tx, tc.givenUserEmail)
				if err != nil {
					require.Equal(t, tc.expectedError, err)
				} else {
					require.Nil(t, err)
					require.Equal(t, tc.expectedResult, result)
				}
			})
		})
	}
}

func TestConnectFriends(t *testing.T) {
	testCases := []struct {
		desc           string
		friendArray    []string
		expectedResult bool
		expectedError  *feature.ResponseError
	}{
		{
			desc:           "Retrieve success",
			friendArray:    []string{"andy@example.com", "dang@example.com"},
			expectedResult: true,
		},
	}

	for _, tc := range testCases {
		db := db.InitDatabase()
		defer db.Close()
		t.Run(tc.desc, func(t *testing.T) {
			var ConnectRequest model.FriendConnectionRequest
			ConnectRequest.Friends = tc.friendArray
			result, err := ConnectFriends(db, ConnectRequest)
			if err != nil {
				require.Error(t, err, tc.expectedError)
			} else {
				require.Nil(t, err)
				require.Equal(t, tc.expectedResult, result.Success)
			}
		})
	}
}

func TestFriendList(t *testing.T) {
	testCase := []struct {
		dest           string
		givenUserEmail string
		expectedResult bool
		expectedError  *feature.ResponseError
	}{
		{
			dest:           "Retrieve success ",
			givenUserEmail: "andy@example.com",
			expectedResult: true,
		},
	}
	for _, tc := range testCase {
		db := db.InitDatabase()
		defer db.Close()
		result, err := FriendList(db, tc.givenUserEmail)
		if err != nil {
			require.Error(t, err, tc.expectedError)
		} else {
			require.Nil(t, err)
			require.Equal(t, tc.expectedResult, result.Success)
		}
	}
}

func TestCommonFriends(t *testing.T) {
	testCase := []struct {
		desc           string
		commonFriends  []string
		expectedResult bool
		expectedError  *feature.ResponseError
	}{
		{
			desc:           "Retrieve success",
			commonFriends:  []string{"andy@example.com", "dang@example.com"},
			expectedResult: true,
		},
	}
	for _, tc := range testCase {
		db := db.InitDatabase()
		defer db.Close()
		var commonRequest model.CommonFriendRequest
		commonRequest.Friends = tc.commonFriends

		result, err := CommonFriends(db, commonRequest)
		if err != nil {
			require.Error(t, err, tc.expectedError)
		} else {
			require.Nil(t, err)
			require.Equal(t, tc.expectedResult, result.Success)
		}
	}
}

func TestSubscription(t *testing.T) {
	testCase := []struct {
		desc             string
		subscribeRequest model.SubscriptionRequest
		expectedResult   bool
		expectedError    *feature.ResponseError
	}{
		{
			desc: "Retrieve success",
			subscribeRequest: model.SubscriptionRequest{
				Requestor: "tu@example.com",
				Target:    "andy@example.com",
			},
			expectedResult: true,
		},
	}
	for _, i := range testCase {
		db := db.InitDatabase()
		defer db.Close()
		result, err := Subscription(db, i.subscribeRequest)
		if err != nil {
			require.Error(t, err, i.expectedError)
		} else {
			require.Nil(t, err)
			require.Equal(t, i.expectedResult, result.Success)
		}
	}
}

func TestBlocked(t *testing.T) {
	testCase := []struct {
		desc           string
		blockedRequest model.SubscriptionRequest
		expectedResult bool
		expectedError  *feature.ResponseError
	}{
		{
			desc: "Retrieve success",
			blockedRequest: model.SubscriptionRequest{
				Requestor: "andy@example.com",
				Target:    "john@example.com",
			},
			expectedResult: true,
		},
	}
	for _, i := range testCase {
		db := db.InitDatabase()
		defer db.Close()
		result, err := Blocked(db, i.blockedRequest)
		if err != nil {
			require.Error(t, err, i.expectedError)
		} else {
			require.Nil(t, err)
			require.Equal(t, i.expectedResult, result.Success)
		}
	}
}

func TestSendUpdate(t *testing.T) {
	testCase := []struct {
		desc           string
		sendRequest    model.SendUpdateRequest
		expectedResult model.SendUpdateResponse
		expectedError  *feature.ResponseError
	}{
		{
			desc: "Retrieve success",
			sendRequest: model.SendUpdateRequest{
				Sender: "andy@example.com",
				Text:   "Hello World! phuc@example.com",
			},
			expectedResult: model.SendUpdateResponse{
				Success:    true,
				Recipients: []string{"phuc@example.com", "tu@example.com", "dang@example.com"},
			},
		},
	}
	for _, i := range testCase {
		db := db.InitDatabase()
		defer db.Close()
		result, err := SendUpdate(db, i.sendRequest)
		if err != nil {
			require.Error(t, err, i.expectedError)
		} else {
			require.Nil(t, err)
			require.Equal(t, i.expectedResult, result)
		}
	}
}
