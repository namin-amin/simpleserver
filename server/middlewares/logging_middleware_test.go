package middlewares

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/namin-amin/simpleserver/server"
	"github.com/stretchr/testify/assert"
)

func TestRequestLoggerMiddleWare(t *testing.T) {
	var expectedRequest http.Request
	tHandler := server.RouteHandler(func(w http.ResponseWriter, r *http.Request) error {
		expectedRequest = *r
		return nil
	})

	tLogger := testLogger{}
	r := RequestLogger(&tLogger)

	tRequest := http.Request{
		Header: http.Header{},
		URL:    &url.URL{Path: "http://test.com"},
		Method: "GET",
	}
	tRequest.Header.Set(server.REQUEST_ID, "123")

	err := r(tHandler)(&testWritter{}, &tRequest)

	assert.Nil(t, err)
	assert.Equal(t, len(tLogger.messages), 2)
	assert.Equal(t, expectedRequest, tRequest)
	assert.Equal(t, tLogger.messages[0], "requestId 123 started")
	assert.Equal(t, tLogger.messages[1], "requestId 123 completed")
}

func TestRequestLoggerMiddleWareWhenNoRequestIdMiddleWareUsed(t *testing.T) {
	var expectedRequest http.Request
	tHandler := server.RouteHandler(func(w http.ResponseWriter, r *http.Request) error {
		expectedRequest = *r
		return nil
	})

	tLogger := testLogger{}
	r := RequestLogger(&tLogger)

	tRequest := http.Request{
		Header: http.Header{},
		URL:    &url.URL{Path: "http://test.com"},
		Method: "GET",
	}

	err := r(tHandler)(&testWritter{}, &tRequest)

	assert.Nil(t, err)
	assert.Equal(t, len(tLogger.messages), 2)
	assert.Equal(t, expectedRequest, tRequest)
	assert.Equal(t, tLogger.messages[0], "requestId "+expectedRequest.Header.Get(server.REQUEST_ID)+" started")
	assert.Equal(t, tLogger.messages[1], "requestId "+expectedRequest.Header.Get(server.REQUEST_ID)+" completed")
}