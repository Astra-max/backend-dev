package handlers

import (
	"net/http"
	"github.com/Astra-max/backend-dev/go/0x00-go/pkg/server"
)

func AuthUser(w http.ResponseWriter, req *http.Request) {
	data := map[string]string{"user": "welcome"}
	server.Json(w, 200, data)
}