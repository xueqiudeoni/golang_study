package main

import (
	"net/http"
	"tcp_chat/cmd"
)

//var conns []*websocket.Conn

//	func handler(w http.ResponseWriter, r *http.Request) {
//		conn, err := upgrader.Upgrade(w, r, nil)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer conn.Close()
//		conns = append(conns, conn)
//		for {
//			m, p, e := conn.ReadMessage()
//			if e != nil {
//				break
//			}
//			for _, conn := range conns {
//				conn.WriteMessage(websocket.TextMessage, []byte("你说的是"+string(p)+"吗？"))
//			}
//			fmt.Println(m, string(p), e)
//		}
//	}

func main() {
	cmd.HUB = cmd.NewHub()
	go cmd.HUB.Run()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cmd.ServeWS(cmd.HUB, w, r)
	})
	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		cmd.Send(cmd.HUB, w, r)
	})
	http.ListenAndServe(":8080", nil)

}
