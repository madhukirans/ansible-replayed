package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/madhukirans/replayed/pkg/types"
	"github.com/madhukirans/replayed/pkg/server"
)

func TestHandler(t *testing.T) {
	config := types.GetReplayedConfig()
	server.InitServer(config)
	w := PerformGetRequest(t,"GET", "/")
	assert.Equal(t, http.StatusOK, w.Code)
}

func PerformGetRequest(t *testing.T, method, path string) *httptest.ResponseRecorder {
	req1, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	rr1 := httptest.NewRecorder()
	handler1 := http.HandlerFunc(server.Handler)
	handler1.ServeHTTP(rr1, req1)

	return w
}
