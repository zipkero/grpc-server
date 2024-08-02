package router

import (
	"net/http"
)

type Router struct {
	routes map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	r := &Router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}
	return r
}

func (r *Router) AddRoute(method, path string, handler http.HandlerFunc) {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]http.HandlerFunc)
	}
	r.routes[method][path] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if r.routes[req.Method] == nil {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if handler, ok := r.routes[req.Method][req.URL.Path]; ok {
		handler(w, req)
		return
	}
	http.NotFound(w, req)
}
