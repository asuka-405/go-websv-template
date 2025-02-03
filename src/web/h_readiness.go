package web

import (
	"net/http"

	"root/src/lib/libresponse"
)

func h_readiness(w http.ResponseWriter, r *http.Request) {
	libresponse.WithHTML(w, http.StatusOK, "<h1>200 OK</h1>")
}
