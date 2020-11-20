package ihttp

import (
	"encoding/json"
	"net/http"
)

// Respond writes JSON as http response
func Respond(w http.ResponseWriter, httpStatusCode int, object interface{}) {
	if object == nil {
		w.WriteHeader(httpStatusCode)
	} else {
		bytes, err := json.Marshal(object)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(httpStatusCode)
		w.Write(bytes)
	}
}
