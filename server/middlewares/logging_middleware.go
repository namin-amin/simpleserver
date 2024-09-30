package middlewares

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/namin-amin/core/server"
)

func RequestLogger(logger slog.Logger) server.MiddlewareHandler {
	return func(next server.RouteHandler) server.RouteHandler {
		return func(w http.ResponseWriter, r *http.Request) error {
			t := time.Now()
			id := r.Header.Get("reqid")
			logger.Info("requestId "+id+" started",
				"path", r.URL.Path,
				"method", r.Method)

			err := next(w, r)

			logger.Info("requestId "+id+" completed",
				"path", r.URL.Path,
				"method", r.Method,
				"timeTaken", time.Since(t))
			return err
		}
	}
}
