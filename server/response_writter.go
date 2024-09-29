package server

import (
	"encoding/json"
	"net/http"
)

func SendString(statuscode int, content string, w http.ResponseWriter) error {
	w.WriteHeader(statuscode)
	_, err := w.Write([]byte(content))
	return err
}

func SendJson(statuscode int, content interface{}, w http.ResponseWriter) error {
	w.WriteHeader(statuscode)
	serializedContent, err := json.Marshal(content)

	if err != nil {
		return err
	}

	_, err = w.Write([]byte(serializedContent))
	return err
}
