package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	Home(response, request)

	got := response.Body.String()
	want := "Hello World"

	if got != want {
		t.Errorf("got %q wannt %q", got, want)
	}
}
