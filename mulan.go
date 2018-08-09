package mulan

import (
	"fmt"
	"net/http"
)

// OnConnect main handler for per request
type OnConnect struct{}

func (oc *OnConnect) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	http.Error(res, "404 page not found", http.StatusNotFound)
}

// Echo for test
func Echo() int {
	return 1
}
