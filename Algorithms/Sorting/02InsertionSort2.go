package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	ar := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}
	insertionSort(ar)
}

func insertionSort(ar []int) {
	l := len(ar)
	for i := 1; i < l; i++ {
		tmp := ar[i]
		for j := i; j >= 0; j-- {
			if j > 0 && ar[j-1] > tmp {
				ar[j] = ar[j-1]
			} else {
				ar[j] = tmp
				break
			}
		}
		fmt.Println(strings.Trim(fmt.Sprint(ar), "[]"))
	}
}
