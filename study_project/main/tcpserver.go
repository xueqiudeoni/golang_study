package main

import (
	"fmt"
	"net"
)
//TCP服务端
//接受请求的方法
func processConn(conn net.Conn){
	var tmp [128]byte
	for{
		n,err:=conn.Read(tmp[:])
		if err!=nil {
			fmt.Println("read from conn failed,err:",err)
		}
		fmt.Println(conn,string(tmp[:n]))
	}
}
func main(){
	//本地端口启动服务
	listener,err:=net.Listen("tcp","127.0.0.1:20000")
	if err!=nil {
		fmt.Println("start server failed",err)
	}
	//for循环监听，等待建立连接
	for{
		conn,err:=listener.Accept()
		if err!=nil {
			fmt.Println("accept failed",err)
		}
		go processConn(conn)
	}
}