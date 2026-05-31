package ws

import "github.com/gorilla/websocket"

var Upgrader = websocket.Upgrade{
	CheckOrigin: func(r *http.Request) bool {
		return true
	}
}
