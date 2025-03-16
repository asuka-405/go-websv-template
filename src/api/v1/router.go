package apiV1

import (
	"github.com/go-chi/chi/v5"
)

func Initialize() chi.Router {
	var router = chi.NewRouter()
	router.Get("/healthz", h_readiness)
	return router
}
