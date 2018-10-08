package mulan

import (
	"fmt"
	"net/http"
)

//
// ──────────────────────────────────────────────────────────────────────── I ──────────
//   :::::: T Y P E S   D E F I N I T I O N : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────────────────────────
//

// ─── HANDLER INTERFACE ──────────────────────────────────────────────────────────

// Handler ...
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

// Router ...
type Router map[Route]RouterHandler

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
	Routes      map[string]Router
}

// ─── SERVER IMPLEMENTIONS ───────────────────────────────────────────────────────

// Server ...
// init a instance
func Server() *HTTPServer {
	// init a server object
	serv := &HTTPServer{}
	serv.Name = "mulan"
	serv.middlewares = &Middlewares{}
	serv.Routes = map[string]Router{}
	serv.Routes["GET"] = Router{}
	serv.Routes["POST"] = Router{}
	serv.Routes["OPTION"] = Router{}
	serv.Routes["DELETE"] = Router{}
	serv.Routes["PUT"] = Router{}

	return serv
}

// Use ...
// register a middleware in serv struct
func (s *HTTPServer) Use(mid Middleware) {
	s.midLen++
	*s.middlewares = append(*s.middlewares, mid)
}

// CallMids ...
func (s *HTTPServer) CallMids(ctx *Ctx) {
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
	m := ctx._Req.Method
	// fmt.Println(ctx._Req.Method)
	p := Route(ctx._Req.URL.Path)
	// fmt.Println(ctx._Req.URL.Path)
	fmt.Println(p, m)
	// fmt.Println(m)
	fmt.Println(s.Routes)
	fmt.Println(s.Routes[m])
	fmt.Println(s.Routes[m][p])
	if s.Routes[m][p] != nil {
		s.Routes[m][p](ctx)
		return
	}
	http.Error(ctx._Res, "not found when muxing", 404)
	return
}

// Imple http methods for Mux

// Get ...
func (s *HTTPServer) Get(r Route, handler RouterHandler) {
	s.RegisterRoute("GET", r, handler)
	// TODO put this in a trie
}

// RegisterRoute ...
func (s *HTTPServer) RegisterRoute(method string, r Route, handler RouterHandler) {
	s.Routes[method][r] = handler
}

// impl http.Handler
func (s *HTTPServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx := new(Ctx)
	ctx._Req = req
	ctx._Res = res
	s.CallMids(ctx)
	s.Mux(ctx)
}

// ─── CTX IMPLEMENTIONS ──────────────────────────────────────────────────────────

// Send ...
func (c *Ctx) Send(interface{}) error {
	return nil
}

// Echo for test
func Echo() int {
	return 1
}
