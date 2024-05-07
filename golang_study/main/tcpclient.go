package main

import(
	"bufio"
	"fmt"
	"net"
	"os"
)
func getInput() string{
	in:=bufio.NewReader(os.Stdin)
	str,_,err:=in.ReadLine()//inReadLine()有3返回值：[]byte bool error,分别为读取到的信息、是否数据太长导致缓存区太长溢出、是否读取失败
	if err!=nil {
		return err.Error()
	}
	return string(str)
}
func main(){
	//与server端建立连接
	conn,err:=net.Dial("tcp","127.0.0.1:20000")
	if err!=nil {
		fmt.Println("dail 127.0.0.1:20000 failed")
		return
	}
	//发送数据
	conn.Write([]byte(getInput()))
	defer conn.Close()
}