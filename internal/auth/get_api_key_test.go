package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	emptyHeaders := http.Header{}
	got, got_err := GetAPIKey(emptyHeaders)
	want, want_err := "", errors.New("no authorization header included")

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
	if !reflect.DeepEqual(want_err, got_err) {
		t.Fatalf("expected: %v, got: %v", want_err, got_err)
	}
}
