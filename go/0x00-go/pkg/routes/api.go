package routes

import (
	"github.com/Astra-max/backend-dev/go/0x00-go/pkg/handlers"
	"github.com/Astra-max/backend-dev/go/0x00-go/pkg/server"
)

func ApiRoutes(r *server.Router) {
	r.GET("/", handlers.AuthUser)
	r.POST("/", handlers.AuthUser)
	r.GET("/health", handlers.Health)
}
