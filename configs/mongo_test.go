package configs

import (
	"testing"
)

func TestGetMongoURI(t *testing.T) {
	mockMongoConfig := MongoConfig{
		User:     "testUser",
		Password: "testPassword",
		Port:     "1234",
	}
	got := mockMongoConfig.GetMongoURI()
	want := "mongodb://testUser:testPassword@localhost:1234"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
