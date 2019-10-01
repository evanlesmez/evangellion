package main

import (
	"evangellion/db"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {
	for {
		// read in a message

		// TODO snag random song at the end of the previous song
		// TODO client change src for img to new img
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// helpful log statement to show connections
	log.Println("Client Connected")
	// err = ws.WriteMessage(1, []byte("<3"))
	// if err != nil {
	// 	log.Println(err)
	// }
	// TODO find out how to return a value from ticker or go function or access ws from another func
	// * Begin animation schedule
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan bool)
	go func() {
	}()
	for {
		select {
		case <-quit:
			return
		case <-ticker.C:
			a, _ := db.SnagAnimation()
			fmt.Println("Snagged a new one from ", a.Artist)
			err = ws.WriteJSON(a)
		}
	}
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", wsEndpoint)
}

func main() {
	fmt.Println("Open localhost:8080 ")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
