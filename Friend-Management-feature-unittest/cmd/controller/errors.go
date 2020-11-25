package controller

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

var (
	ErrMethodNotAllowed = &ErrorResponse{StatusCode: 405, Message: "Method not allowed"}
	ErrNotFound         = &ErrorResponse{StatusCode: 404, Message: "Resource not found"}
	ErrBadRequest       = &ErrorResponse{StatusCode: 400, Message: "Bad request"}
)

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}
func ServerErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: 500,
		Message:    err.Error(),
	}
}

func ServerSuccessRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: 201,
		Message:    err.Error(),
	}
}
