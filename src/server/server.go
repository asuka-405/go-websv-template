package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	libfs "root/src/lib/libfs"

	"github.com/go-chi/chi/v5"
)

func Serve(router chi.Router) {

	static_dir := os.Getenv("SV_STATIC")
	if static_dir == "" {
		static_dir = "static"
	}
	static_dir = "/" + static_dir

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "src", "public"))
	libfs.FileServer(router, static_dir, filesDir)

	webapp_port := os.Getenv("SV_PORT")
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
