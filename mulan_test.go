package mulan

import (
	"testing"
)

// TestEcho first demo unit test
func TestEcho(t *testing.T) {
	got := Echo()
	want := 1
	if got != want {
		t.Errorf("got %d wang %d", got, want)
	}

}

func TestServer(t *testing.T) {
	t.Run("init server instance", func(t *testing.T) {
		serv := Server()
		if "mulan" != serv.Name {
			t.Error()
		}
	})

	t.Run("try listen", func(t *testing.T) {
		serv := Server()
		err := serv.Listen("19990")
		if err != nil {
			t.Error(err)
		}
	})
}
