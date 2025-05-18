package types

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Middleware = func(http.Handler) http.Handler
type HandlerFunc = func(w http.ResponseWriter, h *http.Request)

type RouterMeta struct {
	MountPoint string
	Router     chi.Router
	Middleware []Middleware
	SubRouter  []RouterMeta
}
