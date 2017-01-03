package main

import (
	"fmt"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	fmt.Scan(&n)
	ar := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}
	quickSortInPlace(ar, 0, len(ar)-1)
}

func quickSortInPlace(ar []int, lo, hi int) {
	if lo < hi {
		p := partitionInPlace(ar, lo, hi)
		quickSortInPlace(ar, lo, p-1)
		quickSortInPlace(ar, p+1, hi)
	}
}

// Lomuto partition scheme
func partitionInPlace(ar []int, lo, hi int) int {
	pivot := ar[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if ar[j] <= pivot {
			ar[i], ar[j] = ar[j], ar[i]
			i++
		}
	}
	ar[i], ar[hi] = ar[hi], ar[i]
	fmt.Println(strings.Trim(fmt.Sprint(ar), "[]"))
	return i
}
