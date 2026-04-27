package server

import (
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)
type MiddleWare func(HandleFunc)HandleFunc

type Router struct {
	router map[string]map[string]HandleFunc
	MiddleWares []MiddleWare
}

func NewRouter() *Router {
	return &Router{
		router: make(map[string]map[string]HandleFunc),
	}
}

func (r *Router) Use(m MiddleWare) {
	r.MiddleWares = append(r.MiddleWares, m)
}

func (r *Router) GET(path string, handler HandleFunc, m ...MiddleWare) {
	r.Handle(path,http.MethodGet, handler, m...)
}

func (r *Router) POST(path string, handler HandleFunc) {
	r.Handle(path, http.MethodPost, handler)
}

func (r *Router) Handle(path, method string, handler HandleFunc, m ...MiddleWare) {
	
	if _, exists := r.router[path]; !exists {
		r.router[path] = make(map[string]HandleFunc)
	}
	all := append(r.MiddleWares, m...)

	for i:=len(all)-1; i>=0; i-- {
		handler = all[i](handler)
	}
	r.router[path][method] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path, ok := r.router[req.URL.Path]

	if !ok {
		http.NotFound(w, req)
		return
	}

	handler, allowed := path[req.Method]

	if !allowed {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	handler(w, req)
}