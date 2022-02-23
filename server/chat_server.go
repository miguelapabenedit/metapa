package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = "localhost:8080"

const (
	TypeMessage = 1
	TypeClose   = 8
)

func main() {
	fmt.Println("im a server")
	http.HandleFunc("/ws", connect)
	http.ListenAndServe(addr, nil)
}

var upgrader = websocket.Upgrader{}
var numCons int
var connections map[int]*websocket.Conn = make(map[int]*websocket.Conn)

func connect(rw http.ResponseWriter, r *http.Request) {
	con, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	numCons++
	var conID = numCons
	connections[conID] = con
	for {
		messageType, p, err := con.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if messageType == TypeMessage {
			var msg string
			err := json.Unmarshal(p, &msg)
			if err != nil {
				fmt.Println(err)
				return
			}
			srvMsg := fmt.Sprintf("ConID[%d] sais ->%s\n", conID, msg)
			fmt.Print(srvMsg)
			mssg, err := json.Marshal(srvMsg)
			if err != nil {
				fmt.Println(err)
				return
			}
			for _, v := range connections {
				v.WriteMessage(TypeMessage, mssg)
			}
		} else {
			fmt.Println(connections)
			delete(connections, conID)
			fmt.Println(connections)
		}
	}

}
