package middlewares

import (
	"net/http"
	"testing"

	"github.com/namin-amin/simpleserver/server"
	"github.com/stretchr/testify/assert"
)

func TestRecoverMiddleware(t *testing.T) {
	tHandler := server.RouteHandler(func(w http.ResponseWriter, r *http.Request) error {
		panic("this is a panic")
	})

	tRequest:= http.Request{Header: http.Header{}}
	tRequest.Header.Set(server.REQUEST_ID,"123")

	err:= Recover(tHandler)(&testWritter{}, &tRequest)

	actualError:= err.Error()
	
	assert.Equal(t,"recovering from panic for request with id 123",actualError)
}
