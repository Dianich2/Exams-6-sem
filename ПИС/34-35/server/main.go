package main

import (
	"fmt"
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

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("Upgrade error: ", err)
		return
	}

	defer conn.Close()

	clientName := r.URL.Query().Get("clientName")
	if clientName == "" {
		clientName = "anonymous"
	}

	for {
		messageT, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				log.Println("client " + clientName + " closed")
			} else {
				log.Println("read error: ", err)
			}
			break
		} else {
			log.Println(string(message))
			if err = conn.WriteMessage(messageT, []byte("from server: "+string(message))); err != nil {
				log.Println("write: ", err)
				break
			}
		}

	}
}

func main() {
	http.HandleFunc("/ws", handler)
	fmt.Println("Server running on ", 3000)
	err := http.ListenAndServe("0.0.0.0:3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
