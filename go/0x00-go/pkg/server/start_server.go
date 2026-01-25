package server

import (
	"net/http"
	"log"
)

func Run() {
	r := NewRouter()
	srv := &http.Server{
		Addr: ":8080",
		Handler: r,
	}
	log.Printf("server listening on %s\n", srv.Addr)
	log.Fatalln(srv.ListenAndServe())
}