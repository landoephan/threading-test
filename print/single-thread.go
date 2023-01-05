package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func SingleThreadPrint() {
	slice := MakeRange()
	sliceLength := len(slice)
	var wg sync.WaitGroup
	wg.Add(sliceLength)
	fmt.Println("Running single thread for loop…")
	start := time.Now()
	for i := 0; i < sliceLength; i++ {
		worker(&wg, slice[i])
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("\n------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Finished single thread for loop")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------")
	log.Printf("Loop took %s", elapsed)
}
