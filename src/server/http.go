package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"root/src/api"
	libfs "root/src/lib/libfs"
	"root/src/web"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func initializeHTTPServer(router chi.Router) {

	static_dir := os.Getenv("GWT_SV_STATIC")
	if static_dir == "" {
		static_dir = "static"
	}
	static_dir = "/" + static_dir

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "src", "public"))
	libfs.FileServer(router, static_dir, filesDir)

	webapp_port := os.Getenv("GWT_HTTP_PORT")
	if webapp_port == "" {
		webapp_port = "8080"
	}

	server := &http.Server{
		Handler: router,
		Addr:    ":" + webapp_port,
	}

	fmt.Printf("Server running @ port %v  | address: http://localhost:%v", webapp_port, webapp_port)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("An error occured: %v\n", err)
		log.Fatal("Server Crashed")
	}

}

func BootHTTPServer() {
	api_router := api.Initialize()
	web_router := web.Initialize()

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Mount("/api", api_router)
	router.Mount("/", web_router)

	initializeHTTPServer(router)
}
