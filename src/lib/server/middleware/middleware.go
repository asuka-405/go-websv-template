package middleware

import (
	"fmt"
	"net/http"
	"spotlight/src/lib/types"
	"sync"
)

func Compose(f types.HandlerFunc) types.Middleware {
	return ComposeWithSetup(f, nil, nil)
}

func ComposeWithSetup(f types.HandlerFunc, setup func(args ...any), args ...any) types.Middleware {
	var once sync.Once
	return func(next http.Handler) http.Handler {
		if setup != nil {
			once.Do(func() {
				setup(args...)
			})

		}
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			f(w, r)
			next.ServeHTTP(w, r)
		})
	}
}

// func GetLogger(path string) types.Middleware {
// 	return ComposeWithSetup(fileLogger, SetupLogToFile, path)
// }

// func SetupLogToFile(args ...any) {
// 	path := args[0].(string)

// 	dir := filepath.Dir(path)
// 	if _, err := os.Stat(dir); os.IsNotExist(err) {
// 		_ = os.MkdirAll(dir, 0755)
// 	}

// 	if _, err := os.Stat(path); os.IsNotExist(err) {
// 		_, _ = os.Create(path)
// 	}

// 	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Fatalf("failed to open log file: %v", err)
// 	}

// 	mw := io.MultiWriter(os.Stdout, f)
// 	os.Stdout = f     // overwrite stdout, affects chi middleware.Logger
// 	log.SetOutput(mw) // affects stdlib log
// }

// func fileLogger(w http.ResponseWriter, r *http.Request) {
// 	file, err := os.OpenFile("access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	defer file.Close()

// 	logLine := fmt.Sprintf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL.Path)
// 	if _, err := file.WriteString(logLine); err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// }

var SampleMiddleware = ComposeWithSetup(sampleMiddleware, sampleSetup)

func sampleMiddleware(w http.ResponseWriter, r *http.Request) {
	fmt.Println("runs every time")
}
func sampleSetup(args ...any) {
	fmt.Println("one time setup for sample middleware")
}
