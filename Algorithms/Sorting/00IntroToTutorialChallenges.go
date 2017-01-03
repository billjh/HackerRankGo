package main

import (
	"fmt"
	"sort"
)

func main() {
	var v, n int
	fmt.Scan(&v, &n)
	ar := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}
	// binary search V from ar
	index := sort.SearchInts(ar, v)
	fmt.Println(index)
}
