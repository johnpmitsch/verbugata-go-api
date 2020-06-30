package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestingIndex(t *testing.T) {
	mux := http.NewServeMux()

	// create http.Handler
	handler := IndexHandler()

	// run server using httptest
	server := httptest.NewServer(handler)
	defer server.Close()
}
