package main

import (
	"log"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	args := os.Args
	var clientName string
	if len(args) < 2 {
		clientName = "anonymous"
	} else {
		clientName = args[1]
	}

	url := "ws://localhost:3000/ws?clientName=" + clientName
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err == nil {
		defer conn.Close()
		for i := 1; i <= 5; i++ {
			message := []byte(clientName + " message")
			if err = conn.WriteMessage(websocket.TextMessage, message); err == nil {
				if _, res, err := conn.ReadMessage(); err == nil {
					log.Println(string(res))
				} else {
					log.Println("read: ", err)
					break
				}
			} else {
				log.Println("write: ", err)
				break
			}
			time.Sleep(1 * time.Second)
		}
	} else {
		log.Fatal("dial: ", err)
	}
	conn.WriteMessage(websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "bye"))
}
