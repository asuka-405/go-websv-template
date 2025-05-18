package server

import (
	"compress/gzip"
	"net/http"
	"os"
	"spotlight/src/api"
	"spotlight/src/lib/server"
	"spotlight/src/lib/types"
	"spotlight/src/web"
	"time"

	mw "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

var CORS_DEBUG = true
var HTTP_PUBLIC_DIR = os.Getenv("SPOTLIGHT_PUBLIC_DIR")

func BootSpotlightServer() {
	if HTTP_PUBLIC_DIR == "" {
		HTTP_PUBLIC_DIR = "./src/public"
	}

	subRouters := []types.RouterMeta{
		api.Initialize(),
		web.Initialize(),
	}

	univMiddleware := []types.Middleware{
		mw.RealIP,
		DefaultConsoleLoggerMiddleware(),
		DefaultFileLoggerMiddleware(),
		httprate.LimitByIP(100, time.Minute),
		mw.AllowContentEncoding(
			"identity",
			"gzip",
		),
		mw.Compress(gzip.DefaultCompression),
		mw.CleanPath,
		mw.GetHead,
		mw.Heartbeat("/healthz"),
		mw.Recoverer,
		mw.RequestID,
		mw.Timeout(90 * time.Second),
		cors.Handler(cors.Options{
			ExposedHeaders:     []string{},
			AllowCredentials:   true,
			MaxAge:             300,
			OptionsPassthrough: false,
			Debug:              CORS_DEBUG,
		}),
	}

	var serverMeta = types.RouterMeta{
		Middleware: univMiddleware,
		SubRouter:  subRouters,
	}

	r := server.InitializeRouter(serverMeta)
	router := r.Router
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir(HTTP_PUBLIC_DIR))))
	server.Boot(r)
	// "github.com/go-playground/validator/v10"
}

// look at the following ref before adding a go-chi middleware
// ===================================================================================
// Ref for diff middleware
// this shit will crash the running server
// browsers cant send explicit content-type utf-8 header
// mw.ContentCharset("utf-8"),
// use this bitch to print test request id
// middleware.Compose(func(w http.ResponseWriter, r *http.Request) {
// 	id := mw.GetReqID(r.Context())
// 	log.Printf("RequestID=%s Method=%s URL=%s", id, r.Method, r.URL.Path)
// }),
// use this for web routes
// mw.RedirectSlashes,
// use this for api routes
// mw.StripSlashes
// here just for an example
// use this boy to allow only if
// a header is present
// mw.RouteHeaders
// eg r.With(middleware.RouteHeaders("X-API-KEY", "secret")).Get("/secure", handler)
// use this to add deprication/sunset header to an endpoint
// mw.Sunset
// put ceiling on number of concurrent requests among all users
// mw.Throttle
// to add key-value pair to the context of the request down the request chain
// mw.WithValue
// eg.
// r.Use(middleware.WithValue("userID", 123))

// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
//     userID := r.Context().Value("userID")
//     fmt.Fprintf(w, "User ID: %v", userID)
// })
// take a look at go-chi/stampede for caching
// take a look at go-chi/httpvcr for recording and replaying http requests
