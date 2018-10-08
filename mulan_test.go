package mulan

import (
	"fmt"
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

func TestListen(t *testing.T) {
	serv := Server()
	fmt.Println(serv.Name)
	fmt.Println(serv.Routes)

	serv.Use(func(c *Ctx, next Next) {
		fmt.Println("setting up middles 1")
		next()
	})

	serv.Use(func(c *Ctx, next Next) {
		fmt.Println("setting up middles 2")
		next()
	})

	serv.Use(func(c *Ctx, next Next) {
		fmt.Println("setting up middles 3")
	})

	serv.Get("/hello", func(c *Ctx) {
		fmt.Println("router catch")
		c.Send("hello world")
	})

	fmt.Println(len(*serv.middlewares))

	serv.Listen(":9992")
}
