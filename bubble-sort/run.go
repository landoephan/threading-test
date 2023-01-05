package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func sortingMulti(wg *sync.WaitGroup, arr []int, j int) {
	defer wg.Done()
	if arr[j-1] > arr[j] {
		intermediate := arr[j]
		arr[j] = arr[j-1]
		arr[j-1] = intermediate
	}
}

func sorting(arr []int, j int) {
	if arr[j-1] > arr[j] {
		intermediate := arr[j]
		arr[j] = arr[j-1]
		arr[j-1] = intermediate
	}
}

func singleThread(arr []int) {
	fmt.Println("Running single thread for loop…")
	start := time.Now()
	for i := len(arr); i > 0; i-- {
		for j := 1; j < i; j++ {
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
	for i := len(arr); i > 0; i-- {
		for j := 1; j < i; j++ {
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
	var usedArr = arr5_000
	arr1 := make([]int, len(usedArr))
	copy(arr1, usedArr)
	arr2 := make([]int, len(usedArr))
	copy(arr2, usedArr)

	singleThread(arr1)
	// sorting takes long and doesn't sort
	multiThread(arr2)
}
