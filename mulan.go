package mulan

import (
	"fmt"
	"net/http"
)

// Handler interface
type Handler interface {
	ServeHTTP(res http.ResponseWriter, req *http.Request)
}

// HTTPServer ...
type HTTPServer struct {
	Name       string
	Listenning bool
	handler    Handler
}

// Listen ...
func (s *HTTPServer) Listen(port string) error {
	err := http.ListenAndServe(":"+port, s.handler)

	if err != nil {
		return err
	}

	return nil
}

// Router ...
func (s *HTTPServer) Router() {}

type router struct{}

// Server ...
func Server() *HTTPServer {
	// init a server object
	serv := new(HTTPServer)
	serv.Name = "mulan"
	serv.handler = &OnConnect{}
	fmt.Println(serv)

	return serv
}

// OnConnect main handler for per request
type OnConnect struct{}

func (oc *OnConnect) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	http.Error(res, "not implemented", http.StatusNotFound)
}

// Echo for test
func Echo() int {
	return 1
}
