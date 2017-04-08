package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebsockerManager struct {
}

func NewWebsocketManager() *WebsockerManager {
	return &WebsockerManager{}
}

func (wsm *WebsockerManager) Echo(w http.ResponseWriter, r *http.Request) {

	// do some verifiction for the request
	// Then, upgrade it to websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Can not establish websocket connection")
		return
	}

	go func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			i, b, err := ws.ReadMessage()
			if err != nil {
				log.Print(err)
				return
			}
			err = ws.WriteMessage(i, b)
			if err != nil {
				log.Print(err)
				return
			}
		}
	}(ws)
}
