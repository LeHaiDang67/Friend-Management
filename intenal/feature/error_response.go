package feature

//ResponseError is http error response struct
type ResponseError struct {
	Code        int    `json:"-"`
	Description string `json:"error_description"`
}
