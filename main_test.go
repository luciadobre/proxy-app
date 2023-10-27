package main

import (
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"testing"
)

func TestHandleTodoRequest(t *testing.T) {
	t.Run("get response", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/todos/1", nil)
		rr := httptest.NewRecorder()

		target, _ := url.Parse("https://jsonplaceholder.typicode.com")

		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handleTodoRequest(w, req, target, httputil.NewSingleHostReverseProxy(target))
		}).ServeHTTP(rr, req)

		expected := `{"userId":1,"id":1,"title":"delectus aut autem","completed":false,"foo":"bar"}`
		got := rr.Body.String()

		if rr.Code != http.StatusOK || got != expected {
			t.Errorf("handler returned unexpected result: got %s, want %s", got, expected)
		}
	})
}
