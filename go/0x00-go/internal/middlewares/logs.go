package middlewares

import (
	"net/http"
	"log"
	"github.com/Astra-max/backend-dev/go/0x00-go/pkg/server"
)

func HandleLogs(next server.HandleFunc) server.HandleFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.Method)
		next(w, req)
	}
}