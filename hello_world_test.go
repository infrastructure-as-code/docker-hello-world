package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultRoute(t *testing.T) {
	router := setupRouter("/")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Hello, World!", w.Body.String())
}

func TestHealthRoute(t *testing.T) {
	router := setupRouter("/")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestMetricsRoute(t *testing.T) {
	router := setupRouter("/")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/metrics", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestRoutePrefix(t *testing.T) {
	prefix := "/foobar/"
	router := setupRouter(prefix)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", prefix, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Hello, World!", w.Body.String())
}
