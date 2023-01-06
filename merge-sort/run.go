package main

import (
	"sync"
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

func main() {
	//randomizeSlice(10_000_000)
	var usedArr, _ = readLines("data_100_million")

	arr1 := make([]int, len(usedArr))
	copy(arr1, usedArr)
	_ = sortingSingle(arr1)
	// sorting takes long, but sorts
	//_ = sortingMulti(arr1)
}
