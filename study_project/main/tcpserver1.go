package main

import (
	"fmt"
	"net"
)

func main()  {
	tcpAddr,_:=net.ResolveIPAddr("tcp",":1010")//创建TCP连接
	listener,_:=net.ListenTCP("tcp",tcpAddr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(conn)
	}
}
func handleConnection(conn net.TCPConn)  {
	for{
		buff:=make([]byte,1024)
		n,err:=conn.Read(buff)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(conn.RemoteAddr().String())
		str:="收到了"+string(buff[0:n])
		conn.Write([]byte(str))
	}
}