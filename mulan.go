package mulan

import (
	"fmt"
	"net/http"
)

// ********** Types ***********

// Handler interface
type Handler interface {
	ServeHTTP(res http.ResponseWriter, req *http.Request)
}

// Middleware ...
type Middleware func(c *Ctx, next Next)

// Next ...
type Next func()

// Middlewares ...
type Middlewares []Middleware

// RouterHandler ...
type RouterHandler func(c *Ctx)

// Route ...
// the routing path
type Route string

// Ctx ...
type Ctx struct {
	_Req Request
	_Res Response
}

// Request ...
type Request *http.Request

// Response ...
type Response http.ResponseWriter

// HTTPServer ...
type HTTPServer struct {
	Name        string
	Listenning  bool
	middlewares *Middlewares
	midLen      int64
	Routes      []RouterHandler
}

// ********* implementions *********

// Server ...
// init a instance
func Server() *HTTPServer {
	// init a server object
	serv := new(HTTPServer)
	serv.Name = "mulan"
	serv.middlewares = new(Middlewares)

	return serv
}

// Use ...
// register a middleware in serv struct
func (s *HTTPServer) Use(mid Middleware) {
	s.midLen++
	*s.middlewares = append(*s.middlewares, mid)
}

// SetupMiddlewares ...
func (s *HTTPServer) SetupMiddlewares(ctx *Ctx) {
	NextIter(ctx, 0, s.middlewares)
}

// NextIter ...
func NextIter(ctx *Ctx, c int, m *Middlewares) {
	(*m)[c](ctx, func() {
		c++
		if c < len(*m) {
			NextIter(ctx, c, m)
		}
	})
}

// Listen ...
func (s *HTTPServer) Listen(port string) error {
	fmt.Printf("listening %s\n", port)
	return http.ListenAndServe(port, s)
}

// Mux ...
func (s *HTTPServer) Mux(ctx *Ctx) {
	fmt.Println(ctx._Req.Method)
	fmt.Println(ctx._Req.URL.Path)
	http.Error(ctx._Res, "not found when muxing", 404)
	return
}

// Imple http methods for Mux

// Get ...
func (s *HTTPServer) Get(r Route, handler RouterHandler) {
	// put this in a trie
}

// impl http.Handler
func (s *HTTPServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx := new(Ctx)
	ctx._Req = req
	ctx._Res = res
	s.SetupMiddlewares(ctx)
	s.Mux(ctx)
}

// Echo for test
func Echo() int {
	return 1
}
