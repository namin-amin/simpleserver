package middlewares

import (
	"net/http"

	"github.com/namin-amin/simpleserver/logger"
)

type testLogger struct {
	logger.Logger
	messages []string
}

func (l *testLogger) Info(msg string, v ...any) {
	l.messages = append(l.messages, msg)
}

type testWritter struct {
	HeaderValue int
	content     string
}

func (t *testWritter) Header() http.Header {
	return http.Header{}
}
func (t *testWritter) Write(content []byte) (int, error) {
	t.content = string(content)
	return len(content), nil
}

func (t *testWritter) WriteHeader(statusCode int) {
	t.HeaderValue = statusCode
}
