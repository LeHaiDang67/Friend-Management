package controller

import (
	"bytes"
	"context"
	"errors"
	"friend_management/cmd/mocks"
	"friend_management/intenal/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestConnectFriends(t *testing.T) {
	var jsonStr = []byte(`{"friends":["andy@example","john@example.com"]}`)
	testCases := []struct {
		name         string
		bodyRequest  *bytes.Buffer
		expectedBody string
		mockresponse model.BasicResponse
		mockError    error
	}{
		{
			name:        "retieve success",
			bodyRequest: bytes.NewBuffer(jsonStr),
			mockresponse: model.BasicResponse{
				Success: true,
			},
			expectedBody: "{\"success\":true}\n",
		},
		{
			name:        "retieve failed by CreateFriendConnection error",
			bodyRequest: bytes.NewBuffer(jsonStr),
			mockresponse: model.BasicResponse{
				Success: false,
			},
			mockError:    errors.New("db err"),
			expectedBody: "{\"statusCode\":500,\"message\":\"db err\"}\n",
		},
		{
			name:         "retieve failed by incorrect input",
			bodyRequest:  bytes.NewBuffer([]byte("")),
			expectedBody: "{\"statusCode\":400,\"message\":\"Bad request\"}\n",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/api/friend-management/friends", tt.bodyRequest)
			require.NoError(t, err)
			chiCtx := chi.NewRouter()
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
			rr := httptest.NewRecorder()
			serviceMock := new(mocks.Service)
			serviceMock.On("ConnectFriends", mock.Anything, mock.Anything).Return(tt.mockresponse, tt.mockError)
			handler := ConnectFriends(serviceMock)
			handler.ServeHTTP(rr, req)
			require.Equal(t, tt.expectedBody, rr.Body.String())
		})
	}
}
