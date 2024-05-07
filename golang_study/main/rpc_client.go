package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Req struct {
	NumOne int
	NumTwo int
}
type Res struct {
	Num int
}

func main() {
	req := Req{
		NumOne: 1,
		NumTwo: 2,
	}
	var res Res
	client, err := rpc.DialHTTP("tcp", "localhost:8888")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	//call := client.Call("Server.Add", req, &res)
	call := client.Go("Server.Add", req, &res, nil)
	for {
		select {
		case <-call.Done:
			fmt.Println(res)
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("wait")
		}
	}
}
