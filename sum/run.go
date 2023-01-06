package main

import (
	"sync"
)

var length = 10_000_000_000

func evenNumberSum() {
	sum := 0
	for i := 0; i <= length; i++ {
		if i%2 == 0 {
			sum = sum + i
		}
	}
}

func evenNumberSumMulti(wg *sync.WaitGroup) {
	defer wg.Done()
	evenNumberSum()
}

func oddNumberSum() {
	sum := 0
	for i := 0; i <= length; i++ {
		if i%2 == 1 {
			sum = sum + i
		}
	}
}

func oddNumberSumMulti(wg *sync.WaitGroup) {
	defer wg.Done()
	oddNumberSum()
}

func singleThread() {
	for j := 0; j < 10; j++ {
		evenNumberSum()
		oddNumberSum()
	}
}

func multiThread() {
	var wg sync.WaitGroup
	for j := 0; j < 10; j++ {
		wg.Add(2)
		go evenNumberSumMulti(&wg)
		go oddNumberSumMulti(&wg)
	}
	wg.Wait()
}

func main() {
	singleThread()
	//multiThread()
}
