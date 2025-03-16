package libresponse

import (
	"log"
	"net/http"
)

func WithErr(w http.ResponseWriter, err_code int, message string) {
	if err_code > 499 {
		log.Println("5XX error: ", message)
	}
	type errResponse struct {
		Error string `json:"string"`
	}
	WithJSON(w, err_code, errResponse{message})
}
