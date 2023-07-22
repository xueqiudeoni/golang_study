package main

import (
	"fmt"
	"sync"
	
)
var wg sync.WaitGroup
func putNum(inChan chan int){
	for i:=0;i<10;i++{
		inChan<-(i+1)
	}
	wg.Done()
}
func PrimeNum(inChan chan int,primeChan chan int){
	for value:=range inChan{
		var flag bool=true
		for i:=2;i<value;i++{
			if value%i==0{
				flag=false
				continue
			}
			if flag{
				primeChan<-value
				fmt.Print(flag)
			}
			
		}
	}
	fmt.Println()
	close(primeChan)
	fmt.Println(<-primeChan)
	wg.Done()

}
func PrintPrime(primeChan chan int){
	for value:=range primeChan{
		fmt.Println(value)
	}
	wg.Done()
}
func main(){
	var primeChan chan int
	var inChan chan int
	primeChan=make(chan int,10000)
	inChan=make(chan int,10000)
	wg.Add(1)
	go putNum(inChan)
	for i:=0;i<1;i++{
		wg.Add(1)
		go PrimeNum(inChan,primeChan)
	}
	wg.Add(1)
	PrintPrime(primeChan)
	wg.Wait()
	fmt.Println("主进程执行完毕")
//单向管道的声明
//var ch1 =make(chan<- int,2)仅输入
	var ch1 =make(chan int,2)
	ch1<-10
	ch1<-33
	var ch2 =make(chan int,2)
	ch2<-1
	ch2<-3
	//select多路复用
	for  {
		select{
			case v:=<-ch2:
				fmt.Println("从ch2中读取数据:",v)
			case v:=<-ch1:
				fmt.Println("ch1中读取数据:",v)
		}
	}
}
