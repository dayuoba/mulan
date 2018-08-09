package mulan

import (
	"net/http"
	"testing"
)

// TestEcho first demo unit test
func TestEcho(t *testing.T) {
	got := Echo()
	want := 1
	if got != want {
		t.Errorf("got %d wang %d", got, want)
	}

	ml := &OnConnect{}
	http.ListenAndServe(":9900", ml)
}
