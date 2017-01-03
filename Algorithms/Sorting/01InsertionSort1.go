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
	sortLastEle(ar)
}

func sortLastEle(ar []int) {
	l := len(ar)
	tmp := ar[l-1]
	for i := l - 1; i >= 0; i-- {
		if i > 0 && ar[i-1] > tmp {
			ar[i] = ar[i-1]
		} else {
			ar[i] = tmp
			break
		}
		printArr(ar)
	}
	printArr(ar)
}

func printArr(ar []int) {
	fmt.Println(strings.Trim(fmt.Sprint(ar), "[]"))
}
