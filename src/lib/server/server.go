package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"spotlight/src/lib/types"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func InitializeAndBoot(routerMeta types.RouterMeta) types.RouterMeta {
	routerMeta = InitializeRouter(routerMeta)
	Boot(routerMeta)
	return routerMeta
}

func InitializeRouter(routerMeta types.RouterMeta) types.RouterMeta {
	if routerMeta.Router == nil {
		routerMeta.Router = chi.NewRouter()
	}
	if routerMeta.Middleware != nil {
		routerMeta.Router.Use(routerMeta.Middleware...)
	}
	for _, v := range routerMeta.SubRouter {
		routerMeta.Router.Mount(v.MountPoint, v.Router)
	}
	return routerMeta
}

func Boot(routerMeta types.RouterMeta) {
	portString := os.Getenv("SPOTLIGHT_HTTP_PORT")
	if portString == "" {
		portString = "3000"
	}

	port, _ := strconv.Atoi(portString)

	for {
		server := &http.Server{
			Handler: routerMeta.Router,
			Addr:    fmt.Sprintf(":%d", port),
		}
		fmt.Printf("Server running @ port %d | address: http://localhost:%d\n\n", port, port)

		err := server.ListenAndServe()
		if err != nil && strings.Contains(err.Error(), "address already in use") {
			fmt.Print("\033[1A\033[2K\033[1A\033[2K")
			port++
			continue
		} else if err != nil {
			fmt.Printf("An error occurred: %v\n", err)
			log.Fatal("HTTP Server failed\n")
		}
		break
	}
}
