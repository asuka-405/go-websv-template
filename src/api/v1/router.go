package apiV1

import (
	"net/http"
	"os"
	"path/filepath"

	libfs "github.com/asuka-405/go-webapp/src/lib/fs"
	"github.com/go-chi/chi/v5"
)

func Initialize() chi.Router {

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "src", "public"))

	var router = chi.NewRouter()
	libfs.FileServer(router, "/static", filesDir)
	router.Get("/healthz", h_readiness)
	return router
}
