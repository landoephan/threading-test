package main

import (
	"fmt"
	"sync"
)

func MakeRange() []int {
	a := make([]int, 1000-0)
	for i := range a {
		a[i] = 1 + i
	}
	return a
}

func worker(wg *sync.WaitGroup, val int) {
	defer wg.Done()
	fmt.Printf("%v ", val)
}

func main() {
	SingleThreadPrint()
	MultiThreadPrint()
}
