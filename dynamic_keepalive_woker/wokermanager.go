package main

import "fmt"

type WorkerManager struct {
	workerChan chan *Worker
	numWorkers int
}

func NewWorkerManager(n int) *WorkerManager {
	fmt.Println("new worker manager")
	return &WorkerManager{
		workerChan: make(chan *Worker, n),
		numWorkers: n,
	}
}
func (wm *WorkerManager) StartWorkerPool() {
	fmt.Println("start worker pool")
	for i := 0; i < wm.numWorkers; i++ {
		a := i
		wk := &Worker{
			id: a,
		}
		fmt.Println("worker ", a, " start")
		go wk.work(wm.workerChan)
	}
	wm.KeepLiveWorker()
}
func (wm *WorkerManager) KeepLiveWorker() {
	for worker := range wm.workerChan {
		fmt.Println("worker id:", worker.id, "stopped with err:", worker.err)
		worker.err = nil
		go worker.work(wm.workerChan)
	}
}
