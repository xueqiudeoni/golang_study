package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

// rpc函数要求：函数首字母大写，必须有2个参数，第一个参数是接收的参数，第二个参数是返回给客户端的参数必须是指针类型，函数返回值为error

type Server struct {
}
type Req struct {
	NumOne int
	NumTwo int
}
type Res struct {
	Num int
}

func (s *Server) Add(req Req, res *Res) error {
	time.Sleep(time.Second * 5)
	res.Num = req.NumOne + req.NumTwo
	return nil
}
func main() {
	rpc.Register(new(Server))
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":8888")
	if e != nil {
		fmt.Println("你玩了")
	}
	http.Serve(l, nil)
}
