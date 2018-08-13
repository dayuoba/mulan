package mulan

import (
	"fmt"
	"net/http"
)

// Handler interface
type Handler interface {
	ServeHTTP(res http.ResponseWriter, req *http.Request)
}

// Middleware ...
type Middleware func(c *Context, next Next)

// Next ...
type Next func()

// Middlewares ...
type Middlewares []Middleware

// Router ...
type Router func(c *Context)

// HTTPServer ...
type HTTPServer struct {
	Name        string
	Listenning  bool
	middlewares *Middlewares
	midLen      int64
	Routes      []Router
	Ctx         *Context
}

// Context ...
type Context struct {
	_Request  Request
	_Response Response
}

// Request ...
type Request *http.Request

// Response ...
type Response http.ResponseWriter

func (s *HTTPServer) routes(c *Context) {

}

// Use ...
// register a middleware in serv struct
func (s *HTTPServer) Use(mid Middleware) {
	s.midLen++
	*s.middlewares = append(*s.middlewares, mid)
}

// func (s *HTTPServer) mux(c *Context) Middleware {
// 	return func(c *Context, next Middleware) {

// 	}
// }

// // Next ...
// type Next func()

// Mid ...
type Mid func(c *Context, n Next)

// Mids ...
type Mids []Mid

// SetupMiddlewares ...
func (s *HTTPServer) SetupMiddlewares() {
	NextIter(s.Ctx, 0, s.middlewares)
}

// NextIter ...
func NextIter(ctx *Context, c int, m *Middlewares) {
	(*m)[c](ctx, func() {
		c++
		if c < len(*m) {
			NextIter(ctx, c, m)
		}
	})
}

// Listen ...
func (s *HTTPServer) Listen(port string) error {
	return http.ListenAndServe(":"+port, s)
}

func (s *HTTPServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	s.Ctx._Request = req
	s.Ctx._Response = res
	s.SetupMiddlewares()

	http.Error(res, "not found", 404)
}

// Server ...
func Server() *HTTPServer {
	// init a server object
	serv := new(HTTPServer)
	serv.Name = "mulan"
	serv.middlewares = new(Middlewares)
	fmt.Println(serv)

	return serv
}

// Echo for test
func Echo() int {
	return 1
}
