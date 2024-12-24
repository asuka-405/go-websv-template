package apiV1

import (
	"net/http"

	"github.com/asuka-405/go-webapp/src/lib/libresponse"
)

func h_readiness(w http.ResponseWriter, r *http.Request) {
	libresponse.WithJSON(w, 200, struct{}{})
}
