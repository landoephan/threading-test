package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func sortingSingle(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	n := len(s) / 2
	var l []int
	var r []int
	l = sortingSingle(s[:n])
	r = sortingSingle(s[n:])
	return merge(l, r)
}

func sortingMulti(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	n := len(s) / 2

	wg := sync.WaitGroup{}
	wg.Add(2)

	var l []int
	var r []int

	go func() {
		l = sortingMulti(s[:n])
		wg.Done()
	}()

	go func() {
		r = sortingMulti(s[:n])
		wg.Done()
	}()

	wg.Wait()
	return merge(l, r)
}

func merge(l, r []int) []int {
	ret := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}

func singleThread(arr []int) []int {
	fmt.Println("Running single thread for loop…")
	start := time.Now()
	arr = sortingSingle(arr)
	elapsed := time.Since(start)
	log.Printf("Loop took %s", elapsed)
	fmt.Println("------------------------------------------------------------------------------------------------------------------------")
	return arr
}

func multiThread(arr []int) []int {
	fmt.Println("Running multi thread for loop…")
	start := time.Now()
	arr = sortingMulti(arr)
	elapsed := time.Since(start)
	log.Printf("Loop took %s", elapsed)
	fmt.Println("------------------------------------------------------------------------------------------------------------------------")
	return arr
}

func main() {
	var usedArr = arr100_000
	arr1 := make([]int, len(usedArr))
	copy(arr1, usedArr)
	arr2 := make([]int, len(usedArr))
	copy(arr2, usedArr)

	arr1 = singleThread(arr1)
	// sorting takes long, but sorts
	arr2 = multiThread(arr2)
}
