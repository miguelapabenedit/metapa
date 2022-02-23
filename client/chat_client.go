package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

const (
	TypeMessage = iota + 1
	TypeClose
)

func main() {
	uri := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	c, _, err := websocket.DefaultDialer.Dial(uri.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	num := 1
	for {
		mssg, err := json.Marshal("hello:" + strconv.Itoa(num))
		if err != nil {
			fmt.Println(err)
			return
		}
		num++
		c.WriteMessage(1, mssg)
		if num == 2 {
			c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			time.Sleep(5 * time.Second)
			return
		}
		msgType, p, _ := c.ReadMessage()
		if msgType == TypeMessage {
			if err != nil {
				log.Println(err)
				return
			}
			if msgType == TypeMessage {
				var msg string
				err := json.Unmarshal(p, &msg)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Printf("Server sais->%s\n", msg)
			}
		}
	}
}
