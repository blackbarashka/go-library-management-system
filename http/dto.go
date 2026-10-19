package http

import (
	"encoding/json"
	"time"
)

type CreateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

type ErrorResponse struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func (e *ErrorResponse) ToString() string {
	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return string(b)
}
