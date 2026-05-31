package app

import (
	"github.com/gin-gonic/gin"
	"server-app/internal/handlers"
)

type App struct {
	handle *handlers.Handler
	router *gin.Engine
}

func NewApp() *App {
	return &App{
		handle: handlers.NewHandler(),
		router: gin.New(),
	}
}

func (a *App) routes() {
	a.router.GET("/", a.handle.Home)
	a.router.GET("/ws", a.handle.WsHandler)
}

func (a *App) start() {}

func (a *App) Run() {
	a.router.Run(":8080")
}

