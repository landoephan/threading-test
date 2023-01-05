package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func MultiThreadPrint() {
	slice := MakeRange()
	sliceLength := len(slice)
	var wg sync.WaitGroup
	wg.Add(sliceLength)
	fmt.Println("Running multi thread for loop…")
	start := time.Now()
	for i := 0; i < sliceLength; i++ {
		go worker(&wg, slice[i])
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("\n------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Finished multi thread for loop")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------")
	log.Printf("Loop took %s", elapsed)
}
