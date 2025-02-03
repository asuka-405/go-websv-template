package web

import (
	// apiV1 "root/src/api/v1"
	"log"
	"net/http"

	"root/src/lib/libresponse"
	"root/src/lib/libtemplate"

	"github.com/go-chi/chi/v5"
)

var ViewEngine = libtemplate.NewTemplateEngine("src/web/views")

func Initialize() chi.Router {

	err := ViewEngine.LoadTemplates()
	if err != nil {
		log.Fatal(err)
	}

	head := ViewEngine.RenderWithLogs("default-head.wc", nil)

	data := map[string]string{
		"title": "Hello, World!",
		"head":  head,
		"body":  "Hello, World!",
	}

	rendered := ViewEngine.RenderWithLogs("index.layout.html", data)

	router := chi.NewRouter()
	router.Get("/healthz", h_readiness)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		libresponse.WithHTML(w, http.StatusOK, rendered)
	})

	return router
}
