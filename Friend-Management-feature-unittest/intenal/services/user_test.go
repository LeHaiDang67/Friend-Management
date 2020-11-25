package services

//TestGetUser test GetUser func
// func TestGetUser(t *testing.T) {
// 	testCases := []struct {
// 		desc           string
// 		givenUserEmail string
// 		expectedResult model.User
// 		expectedError  *feature.ResponseError
// 	}{
// 		{
// 			desc:           "Retrieve success",
// 			givenUserEmail: "tom@example.com",
// 		},
// 		{
// 			desc:           "Retrieve error",
// 			givenUserEmail: "andy@example.com",
// 		},
// 	}

// 	for _, i := range testCases {
// 		db := db.InitDatabase()
// 		defer db.Close()
// 		t.Run(i.desc, func(t *testing.T) {
// 			email := i.givenUserEmail
// 			result, err := repo.GetUser(db, email)
// 			if err != nil {
// 				require.Equal(t, i.expectedError, err)
// 			} else {
// 				require.Nil(t, err)
// 				require.Equal(t, i.expectedResult, result)
// 			}
// 		})
// 	}
// }

// func TestConnectFriends(t *testing.T) {
// 	testCases := []struct {
// 		desc           string
// 		friendArray    []string
// 		expectedResult bool
// 		expectedError  *feature.ResponseError
// 	}{
// 		{
// 			desc:           "Retrieve success",
// 			friendArray:    []string{"andy@example.com", "dang@example.com"},
// 			expectedResult: true,
// 		},
// 	}

// 	for _, i := range testCases {
// 		db := db.InitDatabase()
// 		defer db.Close()
// 		t.Run(i.desc, func(t *testing.T) {
// 			var ConnectRequest model.FriendConnectionRequest
// 			ConnectRequest.Friends = i.friendArray
// 			result, err := repo.ConnectFriends(db, ConnectRequest)
// 			if err != nil {
// 				require.Error(t, err, i.expectedError)
// 			} else {
// 				require.Nil(t, err)
// 				require.Equal(t, i.expectedResult, result.Success)
// 			}
// 		})
// 	}
// }

// func TestFriendList(t *testing.T) {
// 	testCase := []struct {
// 		dest           string
// 		givenUserEmail string
// 		expectedResult bool
// 		expectedError  *feature.ResponseError
// 	}{
// 		{
// 			dest:           "Retrieve success ",
// 			givenUserEmail: "andy@example.com",
// 			expectedResult: true,
// 		},
// 	}
// 	for _, i := range testCase {
// 		db := db.InitDatabase()
// 		defer db.Close()
// 		result, err := repo.FriendList(db, i.givenUserEmail)
// 		if err != nil {
// 			require.Error(t, err, i.expectedError)
// 		} else {
// 			require.Nil(t, err)
// 			require.Equal(t, i.expectedResult, result.Success)
// 		}
// 	}
// }

// func TestCommonFriends(t *testing.T) {
// 	testCase := []struct {
// 		dest           string
// 		commonFriends  []string
// 		expectedResult bool
// 		expectedError  *feature.ResponseError
// 	}{
// 		{
// 			dest:           "Retrieve success",
// 			commonFriends:  []string{"andy@example.com", "dang@example.com"},
// 			expectedResult: true,
// 		},
// 	}
// 	for _, i := range testCase {
// 		db := db.InitDatabase()
// 		defer db.Close()
// 		var commonRequest model.CommonFriendRequest
// 		commonRequest.Friends = i.commonFriends

// 		result, err := repo.CommonFriends(db, commonRequest)
// 		if err != nil {
// 			require.Error(t, err, i.expectedError)
// 		} else {
// 			require.Nil(t, err)
// 			require.Equal(t, i.expectedResult, result.Success)
// 		}
// 	}
// }

// func TestSubscription(t *testing.T) {
// 	testCase := []struct {
// 		dest             string
// 		subscribeRequest model.SubscriptionRequest
// 		expectedResult   bool
// 		expectedError    *feature.ResponseError
// 	}{
// 		{
// 			dest: "Retrieve success",
// 			subscribeRequest: model.SubscriptionRequest{
// 				Requestor: "tu@example.com",
// 				Target:    "andy@example.com",
// 			},
// 			expectedResult: true,
// 		},
// 	}
// 	for _, i := range testCase {
// 		db := db.InitDatabase()
// 		defer db.Close()
// 		result, err := repo.Subscription(db, i.subscribeRequest)
// 		if err != nil {
// 			require.Error(t, err, i.expectedError)
// 		} else {
// 			require.Nil(t, err)
// 			require.Equal(t, i.expectedResult, result.Success)
// 		}
// 	}
// }

// func TestBlocked(t *testing.T) {
// 	testCase := []struct {
// 		dest           string
// 		blockedRequest model.SubscriptionRequest
// 		expectedResult bool
// 		expectedError  *feature.ResponseError
// 	}{
// 		{
// 			dest: "Retrieve success",
// 			blockedRequest: model.SubscriptionRequest{
// 				Requestor: "andy@example.com",
// 				Target:    "john@example.com",
// 			},
// 			expectedResult: true,
// 		},
// 	}
// 	for _, i := range testCase {
// 		db := db.InitDatabase()
// 		defer db.Close()
// 		result, err := repo.Blocked(db, i.blockedRequest)
// 		if err != nil {
// 			require.Error(t, err, i.expectedError)
// 		} else {
// 			require.Nil(t, err)
// 			require.Equal(t, i.expectedResult, result.Success)
// 		}
// 	}
// }

// func TestSendUpdate(t *testing.T) {
// 	testCase := []struct {
// 		dest           string
// 		sendRequest    model.SendUpdateRequest
// 		expectedResult model.SendUpdateResponse
// 		expectedError  *feature.ResponseError
// 	}{
// 		{
// 			dest: "Retrieve success",
// 			sendRequest: model.SendUpdateRequest{
// 				Sender: "andy@example.com",
// 				Text:   "Hello World! phuc@example.com",
// 			},
// 			expectedResult: model.SendUpdateResponse{
// 				Success:    true,
// 				Recipients: []string{"phuc@example.com", "tu@example.com", "dang@example.com"},
// 			},
// 		},
// 	}
// 	for _, i := range testCase {
// 		db := db.InitDatabase()
// 		defer db.Close()
// 		result, err := repo.SendUpdate(db, i.sendRequest)
// 		if err != nil {
// 			require.Error(t, err, i.expectedError)
// 		} else {
// 			require.Nil(t, err)
// 			require.Equal(t, i.expectedResult, result)
// 		}

// 	}

// }
