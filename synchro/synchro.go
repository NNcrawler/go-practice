package synchro

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func or(r *http.Request) bool {
	return true
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     or,
}

func Run() {
	http.HandleFunc("/sync", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		for {
			messageType, _, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			
			if err:= conn.WriteMessage(messageType, []byte("3")); err != nil {
				log.Println(err)
				return
			}
		}
	})

	http.ListenAndServe(":1111", nil)
}
