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
	fmt.Println(strings.Trim(fmt.Sprint(partition(ar)), "[]"))
}

func partition(ar []int) []int {
	pivot := ar[0]
	left, right := []int{}, []int{}
	for i := 1; i < len(ar); i++ {
		if ar[i] <= pivot {
			left = append(left, ar[i])
		} else {
			right = append(right, ar[i])
		}
	}
	return append(append(left, pivot), right...)
}
