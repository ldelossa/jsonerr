package jsonerr

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	ID      string `json:"id"`
	Code    string `json:"id"`
	Message string `json:"message"`
}

func NewResponse(id string, code string, message string, args ...interface{}) *Response {
	m := fmt.Sprintf(message, args...)
	return &Response{
		ID:      id,
		Code:    code,
		Message: m,
	}
}

// JsonError works like http.Error but uses our response
// struct as the body of the response. Like http.Error
// you will still need to call a naked return in the http handler
func Error(w http.ResponseWriter, r *Response, httpcode int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(httpcode)
	b, _ := json.Marshal(r)

	w.Write(b)
}
