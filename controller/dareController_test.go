package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// ---------------------- assertion ----------------------
// assertMessage checks whether the return message is the same
func assertResponseBody(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

// assertMessage checks whether the returned statusCode
func assertStatus(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got status %d wnat %d", got, want)
	}
}

// ---------------------- GET ----------------------
func TestHome(t *testing.T) {
	// Instead of using the actual handler Home(w, r)
	// We create a mock server for the test(which handles the Home function)
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)

	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	mux.ServeHTTP(response, request)

	got := response.Body.String()
	want := "{\"Message\":\"welcome to drunk dares\"}"

	assertResponseBody(t, got, want)
	assertStatus(t, response.Code, http.StatusOK)
}

// ---------------------- POST ----------------------
//func mockAddToDatabase(d dare, db dareList) {
//	db = append(db, d)
//}
//
//func TestAddDare(t *testing.T) {
//	request, _ = http.NewRequest(http.MethodPost, "/Dare", nil)
//	response := httptest.NewRecorder()
//
//	AddDare(request, response)
//}
