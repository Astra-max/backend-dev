package main

import (
	"github.com/Astra-max/backend-dev/go/0x00-go/pkg/server"
	"github.com/Astra-max/backend-dev/go/0x00-go/pkg/routes"
)

func main() {
	r := server.NewRouter()
	routes.ApiRoutes(r)
	server.Run(r)
}
