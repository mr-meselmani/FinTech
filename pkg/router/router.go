package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	router *mux.Router
}

func NewRouter(r *mux.Router) *Router {
	// Create and return a new Router instance
	return &Router{
		router: r,
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Implement any middleware logic or custom handling of requests
	// before passing them to the underlying router
	r.router.ServeHTTP(w, req)
}
