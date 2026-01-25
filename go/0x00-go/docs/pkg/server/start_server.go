package server

import (
	"log"
	"net/http"
	"time"
)

func Run(r http.Handler) {

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
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
