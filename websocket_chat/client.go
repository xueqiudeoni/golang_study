package main

import (
	"context"
	"fmt"
	"log"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	conn, _, err := websocket.Dial(ctx, "ws://localhost:2021/ws", nil)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close(websocket.StatusInternalError, "内部错误")
	err = wsjson.Write(ctx, conn, "Hello Websocket Server")
	if err != nil {
		log.Println(err)
	}
	var v interface{}
	err = wsjson.Read(ctx, conn, &v)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("接收到服务器端响应：%v\n", v)
	conn.Close(websocket.StatusNormalClosure, "")
}
