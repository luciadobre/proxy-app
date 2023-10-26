package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	t.Run("return server response", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/todos/1", nil)
		response := httptest.NewRecorder()

		Server(response, request)

		got := response.Body.String()
		want := "hello world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
