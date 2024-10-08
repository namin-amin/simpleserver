package middlewares

import (
	"github.com/namin-amin/simpleserver/logger"
	"github.com/namin-amin/simpleserver/server"
	"net/http"
	"time"
)

func RequestLogger(logger logger.Logger) server.MiddlewareHandler {
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
