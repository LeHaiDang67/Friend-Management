package controller

import (
	"encoding/json"
	"friend_management/internal/db"
	"friend_management/internal/feature"
	"friend_management/internal/feature/model"
	"friend_management/pkg/testutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

//TestGetUser test GetUser func
func TestGetUser(t *testing.T) {
	testCases := []struct {
		desc           string
		givenURL       string
		expectedResult interface{}
		shouldError    bool
	}{
		{
			desc:     "Should get user successfully",
			givenURL: "/friend?email=myemail@gmail.com",
			expectedResult: model.User{
				Email: "myemail@gmail.com",
			},
			shouldError: false,
		},
		{
			desc:     "Should get user fail when missing email param",
			givenURL: "/friend",
			expectedResult: feature.ResponseError{
				Error:       "missing_input_param",
				Description: "Error missing email param",
			},
			shouldError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			testutil.WithTxDB(t, func(tx db.BeginnerExecutor) {
				// Given
				testutil.LoadTestDataFile(t, tx, "cmd/controller/testData/01_get_user.sql")

				// When:
				req := httptest.NewRequest("GET", tc.givenURL, nil)
				w := httptest.NewRecorder()
				GetUser(tx).ServeHTTP(w, req)

				if tc.shouldError {
					var result feature.ResponseError
					json.NewDecoder(w.Body).Decode(&result)
					require.Equal(t, tc.expectedResult, result, "GetUser() Error: "+tc.desc)
				} else {
					var result model.User
					json.NewDecoder(w.Body).Decode(&result)
					require.Equal(t, tc.expectedResult, result, "GetUser() Success: "+tc.desc)
				}
			})
		})
	}
}
