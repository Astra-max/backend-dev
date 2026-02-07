package routes

import (
	"github.com/Astra-max/backend-dev/go/0x00-go/pkg/handlers"
	"github.com/Astra-max/backend-dev/go/0x00-go/pkg/server"
)

func ApiRoutes(r *server.Router) {
	r.GET("/", handlers.AuthUser)
	r.POST("/login", handlers.AuthUser)
	r.POST("/sign-up", handlers.SingUpUser)
	r.POST("/forgot-pass", handlers.ForgotPass)
	r.POST("/refresh", handlers.RefreshToken)
	r.GET("/health", handlers.Health)
}
