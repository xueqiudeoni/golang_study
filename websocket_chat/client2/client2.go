package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"os"
)

func main() {
	dl := websocket.Dialer{}
	conn, _, e := dl.Dial("ws://localhost:8080", nil)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer conn.Close()
	go send(conn)
	for {
		m, p, e := conn.ReadMessage()
		if e != nil {
			break
		}
		fmt.Println(m, string(p))
	}
}
func send(conn *websocket.Conn) {
	for {
		r := bufio.NewReader(os.Stdin)
		l, _, _ := r.ReadLine()
		conn.WriteMessage(websocket.TextMessage, l)
	}
}
