package routes

import (
	"ascii-art-web/internals/handlers"
	"net/http"
)

func ApiRoutes(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("templates"))
	mux.Handle("/templates/", http.StripPrefix("/templates/", fs))

	mux.HandleFunc("/", handlers.ServeTemplate)
	mux.HandleFunc("/ascii-art", handlers.HandleAsciiArt)
}
