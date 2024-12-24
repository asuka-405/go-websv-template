package api

import (
	apiV1 "github.com/asuka-405/go-webapp/src/api/v1"
	"github.com/go-chi/chi/v5"
)

func Initialize() chi.Router {
	routerV1 := apiV1.Initialize()
	router := chi.NewRouter()
	router.Mount("/v1", routerV1)
	return router
}
