package model

//User is an Object
type User struct {
	Email        string   `json:"email"`
	Friends      []string `json:"friends"`
	Subscription []string `json:"subscription"`
	Blocked      []string `json:"blocked"`
}

//BasicResponse return json response
type BasicResponse struct {
	Success bool `json:"success"`
}

//FriendConnectionRequest return friend list
type FriendConnectionRequest struct {
	Friends []string `json:"friends"`
}

//FriendListRequest is...
type FriendListRequest struct {
	Email string `json:"email"`
}

//FriendListResponse is...
type FriendListResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}

//CommonFriendRequest is...
type CommonFriendRequest struct {
	Friends []string `json:"friends"`
}

//SubscriptionRequest is...
type SubscriptionRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

//SendUpdateRequest is...
type SendUpdateRequest struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

//SendUpdateResponse is...
type SendUpdateResponse struct {
	Success    bool     `json:"success"`
	Recipients []string `json:"recipients"`
}
