package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"server-app/internal/services"
)

func (h *Handler) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "it works",
	})
}

func (h *Handler) WsHandler(c *gin.Context) {
	conn, err := services.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer conn.Close()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			break
		}
	}
}

// prevent unused import error
var _ = websocket.Upgrader{}
