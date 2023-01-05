package main

import (
	"fmt"
	"sync"
	"time"
)

func SingleThreadWorker() {
	var wg sync.WaitGroup
	var maxWorkers = 10

	fmt.Println("Running multi thread for loop…")
	start := time.Now()
	for i := 0; i < maxWorkers; i++ {
		fmt.Printf("Adding worker %d \n", i)
		wg.Add(1)
		worker(&wg, i)
	}

	fmt.Printf("\nWaiting for %d workers to finish\n", maxWorkers)
	elapsed := time.Since(start)
	fmt.Println("\n------------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("All Workers Completed in %s\n", elapsed)
	fmt.Println("------------------------------------------------------------------------------------------------------------------------")
}
