package main

import (
	"github.com/madhukirans/replayed/pkg/server"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	router := server.StartServer(nil)
	w := PerformGetRequest(router, "GET", "/")
	assert.Equal(t, http.StatusOK, w.Code)
}

func PerformGetRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
