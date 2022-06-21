package api

import (
	"encoding/json"
	"net/http"
)

// Dare is a type struct for the data within dares(will replace it with a dareDB afterward)
// marshalling for converting query back and forth
type Dare struct {
	DareID       int    `json:"dare_id"`
	DareQuestion string `json:"Dare"`
}

// Home is an interface handler
func Home(w http.ResponseWriter, req *http.Request) {
	// http.Request is a struct with Method, URL...etc.
	// A response writer is used by an HTTP handler to construct an HTTP result
	w.Header().Set("Content-Type", "application/json")
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}

	w.Write([]byte(`"message": "welcome to drunk dares."`))
}

// ListDares is the function of endpoint: GET-drunk/dares for getting all dares
func ListDares(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`"dares": "this is a list of dares"`))
}

// RandomDare is the function of endpoint: GET-drunk/random_dare for getting random dares
func RandomDare(w http.ResponseWriter, r *http.Request) {

}

// CreateDare is the function of endpoint: POST-drunk/dares for creating a new Dare when a Dare doesn't exist
// Authentication required
func CreateDare(w http.ResponseWriter, r *http.Request) {
	// Set the header entries with key to the single element value
	// Header returns the header map that is sent by the WriterHeader
	w.Header().Set("Content-Type", "application/json")
	var newDare Dare
	// Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by v. func (dec *Decoder) Decode(v any) error
	_ = json.NewDecoder(r.Body).Decode(&newDare)
	//InsertDare(newDare)
	//w.Write([]byte(""))
}

// UpdateDare is the function of endpoint: PUT-drunk/Dare for updating dares for the specific id
// Authentication required
func UpdateDare(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

// DeleteDare is the function of endpoint: DELETE-drunk/dares for deleting a specific id
// Authentication required
func DeleteDare(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	// WriteHeader sends an HTTP response header with the provided status code->WriteHeader(statusCode int)
	w.WriteHeader(http.StatusNotFound)

	w.Write([]byte(`{"error": "page not found"}`))
}
