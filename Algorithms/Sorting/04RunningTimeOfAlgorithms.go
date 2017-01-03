package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	ar := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}
	fmt.Println(insertionSort(ar))
}

func insertionSort(ar []int) int {
	shifts := 0
	l := len(ar)
	for i := 1; i < l; i++ {
		tmp := ar[i]
		for j := i; j >= 0; j-- {
			if j > 0 && ar[j-1] > tmp {
				ar[j] = ar[j-1]
				shifts++
			} else {
				ar[j] = tmp
				break
			}
		}
	}
	return shifts
}
