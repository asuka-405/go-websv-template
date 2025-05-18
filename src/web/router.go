package web

import (
	"spotlight/src/lib/server"
	"spotlight/src/lib/types"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Initialize() types.RouterMeta {
	r := chi.NewRouter()
	m := []types.Middleware{
		middleware.AllowContentType(
			"application/json",
			"application/x-www-form-urlencoded",
			"multipart/form-data",
		),
	}
	return server.InitializeRouter(types.RouterMeta{
		MountPoint: "/web",
		Router:     r,
		Middleware: m,
	})
}
