package feature

//ResponseError is http error response struct
type ResponseError struct {
	Error       string `json:"error"`
	Description string `json:"error_description"`
}
