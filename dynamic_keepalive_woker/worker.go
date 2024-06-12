package main

import (
	"fmt"
	"runtime"
	"time"
)

type Worker struct {
	id  int
	err error
}

func (w *Worker) work(workerChan chan<- *Worker) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				w.err = err
			} else {
				w.err = fmt.Errorf("panic happened, worker %d err:%s", w.id, r)
			}
		}
		workerChan <- w
	}()
	fmt.Println("start worker ", w.id)
	time.Sleep(1 * time.Second)
	//panic("worker panic")
	runtime.Goexit()
	return
}
