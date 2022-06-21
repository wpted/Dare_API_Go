package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(Home)
	handler.ServeHTTP(res, req)

	if http.StatusOK != res.Code {

	}

	if `"message": "welcome to drunk dares."` != res.Body.String() {
		t.Error("expected message: welcome to drunk dares got", res.Body.String())
	}
}
