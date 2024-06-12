package main

func main() {
	i := 10
	workemg := NewWorkerManager(i)
	workemg.StartWorkerPool()

}
