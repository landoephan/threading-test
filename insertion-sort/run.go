package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func sorting(arr []int, j int) {
	arr[j], arr[j-1] = arr[j-1], arr[j]
}

func sortingMulti(wg *sync.WaitGroup, arr []int, j int) {
	defer wg.Done()
	arr[j], arr[j-1] = arr[j-1], arr[j]
}

func singleThread(arr []int) {
	fmt.Println("Running single thread for loop…")
	start := time.Now()
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			sorting(arr, j)
		}
	}
	elapsed := time.Since(start)
	log.Printf("Loop took %s", elapsed)
	fmt.Println("------------------------------------------------------------------------------------------------------------------------")
}

func multiThread(arr []int) {
	var wg sync.WaitGroup
	fmt.Println("Running multi thread for loop…")
	start := time.Now()
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			wg.Add(1)
			go sortingMulti(&wg, arr, j)
		}
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("Loop took %s", elapsed)
	fmt.Println("------------------------------------------------------------------------------------------------------------------------")
}

func main() {
	var usedArr = arr100_000
	arr1 := make([]int, len(usedArr))
	copy(arr1, usedArr)
	arr2 := make([]int, len(usedArr))
	copy(arr2, usedArr)

	singleThread(arr1)
	// sorting is faster, but doesn't sort
	multiThread(arr2)
}
