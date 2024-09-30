package middlewares

import (
	"errors"
	"fmt"
	"github.com/namin-amin/simpleserver/server"
	"net/http"
)

func Recover(next server.RouteHandler) server.RouteHandler {

	return func(w http.ResponseWriter, r *http.Request) (err error) {
		id := r.Header.Get("reqid")
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recovered from panic for requestid "+id+"\n", r)
				err = errors.New("recovering from panic for request with id " + id)
			}
		}()

		err = next(w, r)
		return err
	}

}
