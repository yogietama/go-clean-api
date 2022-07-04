package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) GET(url string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(url, f)

}
func (*chiRouter) POST(url string, f func(w http.ResponseWriter, r *http.Request)) {

	chiDispatcher.Post(url, f)
}
func (*chiRouter) SERVE(port string) {
	fmt.Println("Chi HTTP server running on port", port)
	http.ListenAndServe(port, chiDispatcher)

}
