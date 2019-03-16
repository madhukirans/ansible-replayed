package main

import (
	"github.com/madhukirans/replayed/pkg/server"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


func TestHandler(t *testing.T) {
	tt := []struct {
		name  string
		value string
	}{
		{name: "test1", value: "abcdef"},
		{name: "test2", value: ""},
		{name: "test3", value: "x"},
	}

	router := server.StartServer(nil)
	var body string
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			w := PerformPostRequest(router, "POST", "/", strings.NewReader(tc.value))
			assert.Equal(t, http.StatusOK, w.Code)

			w = PerformGetRequest(router, "GET", "/")
			assert.Equal(t, http.StatusOK, w.Code)
			body = body + tc.value
			assert.Equal(t, w.Body.String(), body)
		})
	}
}

func BenchmarkHttp(b *testing.B) {
	router := server.StartServer(nil)
	for i := 0; i < b.N; i++ {
		PerformPostRequest(router, "POST", "/", strings.NewReader("some data string"))
		PerformGetRequest(router, "GET", "/")
	}
}

func PerformPostRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func PerformGetRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}