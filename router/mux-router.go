package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(url string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(url, f).Methods("GET")

}
func (*muxRouter) POST(url string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(url, f).Methods("POST")

}
func (*muxRouter) SERVE(port string) {
	fmt.Println("Mux HTTP server running on port", port)
	http.ListenAndServe(port, muxDispatcher)

}
