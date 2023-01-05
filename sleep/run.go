package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	fmt.Printf("Worker %v: Started ", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %v: Finished ", id)
}

func main() {
	SingleThreadWorker()
	MultiThreadWorker()
}
