package server

import (
	"io"
	"log"
	"net/http"
)

func getReqBody(r *http.Request) string {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("[ERROR] failed to read body: %v", err)

		r.Body.Close()
		return ""
	}

	defer r.Body.Close()

	return string(body)
}

func getReqHeaders(r *http.Request) map[string]string {
	headers := make(map[string]string)
	for name, values := range r.Header {
		headers[name] = values[0]
	}

	return headers
}
