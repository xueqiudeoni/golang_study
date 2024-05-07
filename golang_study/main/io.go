package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func file()  {
	d,err:=os.ReadDir("./")
	fmt.Println(err)
	for _,v:=range d  {
		fmt.Println(v.Info())
		fmt.Println(v.IsDir())
		fmt.Println(v.Name())
		fmt.Println(v.Type())
	}
}
func main(){
	f,err:=os.OpenFile("./test.txt",os.O_CREATE|os.O_RDWR,777)
	if err!=nil{
		fmt.Println(err)
		return
	}
	//for  {
	//	b:=make([]byte,12)
	//	n,err:=f.Read(b)
	//	if err!=nil {
	//		fmt.Println(n,err)
	//		return
	//	}
	//	fmt.Println(string(b),n)
	//}

	f.Seek(0,io.SeekStart)
	n1,err:=f.Write([]byte("123456"))
	if err!=nil {
		fmt.Println(err,n1)
		return
	}
	resder:=bufio.NewReader(f)
	for  {
		str,isPrefix,err:=resder.ReadLine()
		if err!=nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(str),isPrefix)
	}
	f.Close()
}