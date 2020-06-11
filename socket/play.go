package socket

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
	c := make(chan []byte, 10)
	mt := make(chan int, 10)
	conns := make([]*websocket.Conn, 0)

	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		conns = append(conns, conn)

		go func() {
			for {
				messageType, p, err := conn.ReadMessage()
				if err != nil {
					log.Println(err)
					return
				}
				c <- p
				mt <- messageType
			}
		}()

		for {
			message := <-c
			mType := <-mt
			log.Println("received: ", string(message))
			for _, wConn := range conns {
				log.Println(wConn.RemoteAddr)
				if err := wConn.WriteMessage(mType, message); err != nil {
					log.Println(err)
					return
				}
			}
		}
	})

	http.ListenAndServe(":1111", nil)
}
