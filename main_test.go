package main

import (
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"strings"
	"testing"
)

func TestHandleTodoRequest(t *testing.T) {
	t.Run("get response", func(t *testing.T) {
		request, _ := http.NewRequest("GET", "/todos/1", nil)
		reply := httptest.NewRecorder()

		target, _ := url.Parse("https://jsonplaceholder.typicode.com")
		pretendServer := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handleRequest(w, request, target, httputil.NewSingleHostReverseProxy(target))
		})
		pretendServer.ServeHTTP(reply, request)

		expected := `{"userId":1,"id":1,"title":"delectus aut autem","completed":false,"foo":"bar"}`
		got := strings.TrimSpace(reply.Body.String())

		if got != expected {
			t.Errorf("Got %s, want %s", got, expected)
		}
	})
}
