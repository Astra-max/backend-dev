package server

import (
	"log"
	"net/http"
	"time"
)

func Run(r *Router) {
	handler := CORS(r)
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	log.Printf("server listening on %s\n", srv.Addr)
	err := srv.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}
