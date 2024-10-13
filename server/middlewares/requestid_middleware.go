package middlewares

import (
	"net/http"

	"github.com/namin-amin/simpleserver/server"

	"github.com/google/uuid"
)

func RequestId() server.MiddlewareHandler {
	return func(next server.RouteHandler) server.RouteHandler {
		return func(w http.ResponseWriter, r *http.Request) error {
			id := uuid.New()
			r.Header.Add(server.REQUEST_ID, id.String())
			return next(w, r)
		}
	}
}
