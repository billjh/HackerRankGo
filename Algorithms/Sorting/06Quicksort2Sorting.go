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
	quickSort(ar)
}

func quickSort(ar []int) []int {
	if len(ar) <= 1 {
		return ar
	}
	pivot := ar[0]
	left, right := []int{}, []int{}
	// partition
	for i := 1; i < len(ar); i++ {
		if ar[i] <= pivot {
			left = append(left, ar[i])
		} else {
			right = append(right, ar[i])
		}
	}
	// resursively sort sub-arrays
	sorted := append(append(quickSort(left), pivot), quickSort(right)...)
	fmt.Println(strings.Trim(fmt.Sprint(sorted), "[]"))
	return sorted
}
