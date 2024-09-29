package middlewares

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/namin-amin/core/server"
)

func RequestId() server.MiddlewareHandler {
	return func(next server.RouteHandler) server.RouteHandler {
		return func(w http.ResponseWriter, r *http.Request) error {
			id:= uuid.New()
			r.Header.Add("reqid",id.String())
			return next(w,r)
		}
	}
}