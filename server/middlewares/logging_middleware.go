package middlewares

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/namin-amin/simpleserver/logger"
	"github.com/namin-amin/simpleserver/server"
)

func RequestLogger(logger logger.Logger) server.MiddlewareHandler {
	return func(next server.RouteHandler) server.RouteHandler {
		return func(w http.ResponseWriter, r *http.Request) error {
			t := time.Now()
			id := r.Header.Get(server.REQUEST_ID)
			if id == "" {
				id= uuid.NewString()
				r.Header.Set(server.REQUEST_ID,id)
			}

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
