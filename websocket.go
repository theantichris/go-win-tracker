package poker

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type playerServerWs struct {
	*websocket.Conn
}

func (w *playerServerWs) WaitForMsg() string {
	_, msg, err := w.ReadMessage()

	if err != nil {
		log.Printf("error reading from Websocket %v\n", err)
	}

	return string(msg)
}

func (w *playerServerWs) Write(p []byte) (n int, err error) {
	err = w.WriteMessage(1, p)

	if err != nil {
		return 0, err
	}

	return len(p), nil
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func newPlayerServerWs(w http.ResponseWriter, r *http.Request) *playerServerWs {
	conn, err := wsUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("problem upgrading connection to Websockets %v", err)
	}

	return &playerServerWs{conn}
}
