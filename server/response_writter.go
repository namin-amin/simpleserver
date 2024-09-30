package server

import (
	"encoding/json"
	"net/http"
)

func SendString(statusCode int, content string, w http.ResponseWriter) error {
	w.WriteHeader(statusCode)
	_, err := w.Write([]byte(content))
	return err
}

func SendJson(statusCode int, content interface{}, w http.ResponseWriter) error {
	w.WriteHeader(statusCode)
	serializedContent, err := json.Marshal(content)

	if err != nil {
		return err
	}

	_, err = w.Write([]byte(serializedContent))
	return err
}
