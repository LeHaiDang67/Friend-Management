package mocks

import (
	"friend_management/intenal/model"

	"github.com/stretchr/testify/mock"
)

//Service is...
type Service struct {
	mock.Mock
}

// ConnectFriends is...
func (st *Service) ConnectFriends(req model.FriendConnectionRequest) (model.BasicResponse, error) {
	returnVals := st.Called(req)
	r0 := returnVals.Get(0).(model.BasicResponse)
	var r1 error
	if returnVals.Get(1) != nil {
		r1 = returnVals.Get(1).(error)
	}
	return r0, r1
}
