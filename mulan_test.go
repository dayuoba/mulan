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
}

func TestNextIter(t *testing.T) {
	// var mids *Mids
	// mids := make(*Mids, 3)
	mids := new(Mids)
	*mids = append(*mids, func(p int, n Next) {
		println(p)
		n()
	})
	*mids = append(*mids, func(p int, n Next) {
		println(p + 1)
		n()
	})
	*mids = append(*mids, func(p int, n Next) {
		println(p + 2)
		n()
	})
	NextIter(1, 0, mids)
	// t.Error("foo")
}

func BenchmarkNextIter100(b *testing.B) {
	mids := new(Mids)
	for i := 0; i < 100; i++ {
		*mids = append(*mids, func(p int, n Next) {
			// println(p)
			n()
		})
	}

	for n := 0; n < b.N; n++ {
		NextIter(1, 0, mids)
	}
}

func BenchmarkNormalNextIter100(b *testing.B) {
	var mids Mids
	for i := 0; i < 100; i++ {
		mids = append(mids, func(p int, n Next) {
			// println(p)
			n()
		})
	}

	for n := 0; n < b.N; n++ {
		NormalNextIter(1, 0, mids)
	}
}
