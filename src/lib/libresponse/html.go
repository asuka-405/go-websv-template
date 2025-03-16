package libresponse

import (
	"log"
	"net/http"
)

func WithHTML(w http.ResponseWriter, code int, htmlContent string) {
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(code)
	_, err := w.Write([]byte(htmlContent))
	if err != nil {
		log.Printf("Failed to write HTML response, %v", err)
		WithErr(w, 500, err.Error())
	}
}
