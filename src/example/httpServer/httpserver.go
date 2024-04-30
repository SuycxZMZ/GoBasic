package httpServer

import (
	"fmt"
	"net/http"
)

// func ListenAndServe(addr string, handler Handler) error {

// }

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

func PalyerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
