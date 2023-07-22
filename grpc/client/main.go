package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/pb/person"
	"sync"
	"time"
)

func main() {
	l, _ := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	client := person.NewSearchServiceClient(l)
	//res, err := client.Search(context.Background(), &person.PersonReq{
	//	Name: "我是swag"})
	//c, _ := client.SearchIn(context.Background())
	//i := 0
	//for {
	//	if i > 10 {
	//		res, _ := c.CloseAndRecv()
	//		fmt.Println(res)
	//		break
	//	}
	//	time.Sleep(1 * time.Second)
	//	c.Send(&person.PersonReq{Name: "我是进来的信息"})
	//	i++
	//}
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res)
	//c, _ := client.SearchOut(context.Background(), &person.PersonReq{Name: "swag"})
	//for {
	//	res, err := c.Recv()
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println(res)
	//}
	c, _ := client.SearchIO(context.Background())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for {
			time.Sleep(time.Second * 1)
			err := c.Send(&person.PersonReq{Name: "swag"})
			if err != nil {
				wg.Done()
				break
			}
		}
	}()
	go func() {
		for {
			res, err := c.Recv()
			if err != nil {
				fmt.Println(err)
				wg.Done()
				break
			}
			fmt.Println(res)
		}
	}()
	wg.Wait()
}
