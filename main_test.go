package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETRequest(t *testing.T) {
	t.Run("returns request", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
