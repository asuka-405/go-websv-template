package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"spotlight/src/lib/types"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxKeyLogger struct{}

type LoggerOptions struct {
	LogMethod  bool
	LogPath    bool
	LogRemote  bool
	LogHeaders bool
}

var DefaultLoggerOptions = LoggerOptions{
	LogMethod:  true,
	LogPath:    true,
	LogRemote:  true,
	LogHeaders: true,
}

func newFileLogger() *zap.Logger {
	file, err := os.OpenFile("log/server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	ws := zapcore.AddSync(file)

	encCfg := zap.NewProductionEncoderConfig()
	encCfg.TimeKey = "timestamp"
	encCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encCfg), ws, zapcore.InfoLevel)
	return zap.New(core)
}

func newConsoleLogger() *zap.Logger {
	encCfg := zap.NewDevelopmentEncoderConfig()
	encCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encCfg),
		zapcore.Lock(os.Stderr),
		zapcore.DebugLevel,
	)

	return zap.New(core)
}

func DefaultFileLoggerMiddleware() types.Middleware {
	return FileLoggerMiddleware(DefaultLoggerOptions)
}

func DefaultConsoleLoggerMiddleware() types.Middleware {
	return ConsoleLoggerMiddleware(DefaultLoggerOptions)
}

func FileLoggerMiddleware(opts LoggerOptions) types.Middleware {
	return genericLoggerMiddleware(newFileLogger(), opts)
}

func ConsoleLoggerMiddleware(opts LoggerOptions) types.Middleware {
	return genericLoggerMiddleware(newConsoleLogger(), opts)
}

func genericLoggerMiddleware(logger *zap.Logger, opts LoggerOptions) types.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			fields := []zap.Field{}

			if opts.LogMethod {
				fields = append(fields, zap.String("method", r.Method))
			}
			if opts.LogPath {
				fields = append(fields, zap.String("path", r.URL.Path))
			}
			if opts.LogRemote {
				fields = append(fields, zap.String("remote", r.RemoteAddr))
			}
			if opts.LogHeaders {
				for k, v := range r.Header {
					if len(v) > 0 {
						fields = append(fields, zap.String("header."+k, v[0]))
					}
				}
			}

			reqLogger := logger.With(fields...)
			ctx := context.WithValue(r.Context(), ctxKeyLogger{}, reqLogger)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)

			reqLogger.Info("request complete", zap.Duration("duration", time.Since(start)))
		})
	}
}

func LoggerFromContext(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(ctxKeyLogger{}).(*zap.Logger); ok {
		return logger
	}
	return zap.NewNop()
}
