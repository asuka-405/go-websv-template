package api

import (
	"net/http"
	"spotlight/src/lib/types"

	"github.com/go-chi/chi/v5"
)

func Initialize() types.RouterMeta {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte{})
	})

	return types.RouterMeta{
		MountPoint: "/api",
		Router:     r,
	}
}
