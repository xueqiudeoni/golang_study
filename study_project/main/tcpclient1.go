package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main()  {
	tcpAddr,_:=net.ResolveTCPAddr("tcp",":1010")
	conn,err:=net.DialTCP("tcp",nil,tcpAddr)
	if err!=nil{
		fmt.Println(err)
		return
	}
	reader:=bufio.NewReader(os.Stdin)
	for{
		bytes,_,_:=reader.ReadLine()
		conn.Write(bytes)
		rb:=make([]byte,1024)
		rn,_:=conn.Read(rb)
		fmt.Println(rb[0:rn])
	}

}

