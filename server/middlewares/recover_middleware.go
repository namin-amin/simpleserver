package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/namin-amin/simpleserver/server"
)

func Recover(next server.RouteHandler) server.RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) (err error) {
		id := r.Header.Get(server.REQUEST_ID)
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
