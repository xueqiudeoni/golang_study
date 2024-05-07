package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type client struct {
	id      int
	integer int
}
type data struct {
	job    client
	square int
}

var (
	size, _ = strconv.Atoi(os.Args[3])
	clients = make(chan client, size)
	datas   = make(chan data, size)
)

func worker(w *sync.WaitGroup) {
	for cl := range clients {
		sq := cl.integer * cl.integer
		output := data{cl, sq}
		datas <- output
		time.Sleep(time.Second)
	}
	w.Done()
}
func makeWK(n int) {
	var w sync.WaitGroup
	for i := 0; i < n; i++ {
		w.Add(1)
		go worker(&w)
	}
	w.Wait()
	close(datas)
}
func create(n int) {
	for i := 0; i < n; i++ {
		c := client{
			id:      i,
			integer: i,
		}
		clients <- c
	}
	close(clients)
}
func main() {
	fmt.Println("cap(clients):", cap(clients))
	fmt.Println("cap(datas):", cap(datas))
	if len(os.Args) != 3 {
		fmt.Println("need jobs and workers")
	}
	njobs, _ := strconv.Atoi(os.Args[1])
	nworker, _ := strconv.Atoi(os.Args[2])
	go create(njobs)
	makeWK(nworker)
	finished := make(chan bool)
	go func() {
		for i := range datas {
			fmt.Printf("client id:%d\t int:%d\t square:%d\n", i.job.id, i.job.integer, i.square)
		}
		finished <- true
	}()
	fmt.Printf("finished: %v", <-finished)
}
