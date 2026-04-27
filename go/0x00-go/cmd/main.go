package main

import (
	"github.com/Astra-max/backend-dev/go/0x00-go/internal/server"
	"github.com/Astra-max/backend-dev/go/0x00-go/internal/routes"
)

func main() {
	r := server.NewRouter()
	routes.ApiRoutes(r)
	server.Run(r)
}
