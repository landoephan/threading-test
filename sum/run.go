package main

import (
	"fmt"
	"log"
	"sync"
	"time"
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
	fmt.Println("Running single thread for loop…")
	start := time.Now()
	for i := 0; i < 5; i++ {
		evenNumberSum()
		oddNumberSum()
	}
	elapsed := time.Since(start)
	log.Printf("Loop took %s", elapsed)
	fmt.Println("------------------------------------------------------------------------------------------------------------------------")
}

func multiThread() {
	var wg sync.WaitGroup
	fmt.Println("Running multi thread for loop…")
	start := time.Now()
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go evenNumberSumMulti(&wg)
		go oddNumberSumMulti(&wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("Loop took %s", elapsed)
	fmt.Println("------------------------------------------------------------------------------------------------------------------------")
}

func main() {
	singleThread()
	multiThread()
}
