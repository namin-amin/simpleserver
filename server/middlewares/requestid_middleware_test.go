package middlewares

import (
	"net/http"
	"testing"

	"github.com/namin-amin/simpleserver/server"
	"github.com/stretchr/testify/assert"
)

func TestRequestIdMiddleware(t *testing.T) {
	var expectedRequest http.Request
	tHandler := server.RouteHandler(func(w http.ResponseWriter, r *http.Request) error {
		expectedRequest = *r
		return nil
	})

	idMiddleware:= RequestId()
	err:=idMiddleware(tHandler)(&testWritter{},&http.Request{
		Header: http.Header{},
	})

	assert.Nil(t,err)
	assert.NotEmpty(t,expectedRequest.Header.Get(server.REQUEST_ID))
}
