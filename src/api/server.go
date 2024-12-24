package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func Serve(router chi.Router) {

	webapp_port := os.Getenv("PORT")

	server := &http.Server{
		Handler: router,
		Addr:    ":" + webapp_port,
	}

	fmt.Printf("Server running @ port %v\n", webapp_port)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("An error occured: %v\n", err)
		log.Fatal("Server Crashed")
	}

}
