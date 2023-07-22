package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"grpc/pb/person"
	"net"
	"net/http"
	"sync"
)

type personServe struct {
	person.UnimplementedSearchServiceServer
}

func (*personServe) Search(ctx context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	name := req.Name
	res := &person.PersonRes{Name: "我收到了" + name + "的信息,来自grpcgateway", Age: req.Age, Mbody: req.Mbody}
	return res, nil
}

//	func (*personServe) SearchIn(server person.SearchService_SearchInServer) error {
//		for {
//			req, err := server.Recv()
//			if err != nil {
//				server.SendAndClose(&person.PersonRes{Name: "完成了"})
//				break
//			}
//			fmt.Println(req)
//		}
//		return nil
//	}
//
//	func (*personServe) SearchOut(req *person.PersonReq, server person.SearchService_SearchOutServer) error {
//		name := req.Name
//		i := 0
//		for {
//			if i > 5 {
//				break
//			}
//			time.Sleep(time.Second * 1)
//			server.Send(&person.PersonRes{Name: "我拿到了" + name})
//			i++
//		}
//		return nil
//	}
//
// func (*personServe) SearchIO(server person.SearchService_SearchIOServer) error {
//
//		i := 0
//		str := make(chan string)
//		go func() {
//			for {
//				i++
//				req, _ := server.Recv()
//				if i > 10 {
//					str <- "结束"
//					break
//				}
//				str <- req.Name
//			}
//		}()
//		for {
//			s := <-str
//			if s == "结束" {
//				break
//			}
//			server.Send(&person.PersonRes{Name: s})
//		}
//		return nil
//	}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go registerGateway(&wg)
	go registerGrpc(&wg)
	wg.Wait()
}
func registerGateway(wg *sync.WaitGroup) {
	conn, _ := grpc.DialContext(context.Background(), "127.0.0.1:8888", grpc.WithBlock(), grpc.WithInsecure())

	mux := runtime.NewServeMux() //对外开放的mux
	gwServer := http.Server{
		Addr:    ":8090",
		Handler: mux,
	}
	_ = person.RegisterSearchServiceHandler(context.Background(), mux, conn)
	gwServer.ListenAndServe()
	wg.Done()
}
func registerGrpc(wg *sync.WaitGroup) {
	l, _ := net.Listen("tcp", ":8888")
	s := grpc.NewServer()
	person.RegisterSearchServiceServer(s, &personServe{})
	s.Serve(l)
	wg.Done()
}
