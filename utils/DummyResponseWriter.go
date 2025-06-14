package utils

import (
	"log"
	"net/http"
)

type DummyResponseWriter struct {
}

func (d DummyResponseWriter) Header() http.Header {
	return http.Header{}
}

func (d DummyResponseWriter) Write(b []byte) (int, error) {
	log.Printf("Response: %s", b)
	return len(b), nil
}

func (d DummyResponseWriter) WriteHeader(statusCode int) {
	log.Printf("Response status code: %d", statusCode)
}

func (d DummyResponseWriter) Flush() {
}
