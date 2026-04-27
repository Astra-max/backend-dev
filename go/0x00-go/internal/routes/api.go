package routes

import (
	"github.com/Astra-max/backend-dev/go/0x00-go/internal/handlers"
	"github.com/Astra-max/backend-dev/go/0x00-go/internal/server"
)

func ApiRoutes(r *server.Router) {
	r.GET("/", handlers.HomeHandler)
	r.POST("/", handlers.AuthUser)
	r.GET("/health", handlers.Health)
}
