// go web example -websockets

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
}

func main() {

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {

		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// read messages from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			// print msg to console
			fmt.Printf("\n%s sent %s", conn.RemoteAddr(), string(msg))

			// write msg back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}

		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.ListenAndServe(":8000", nil)
}
