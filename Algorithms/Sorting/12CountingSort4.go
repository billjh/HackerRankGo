package main

import (
	"fmt"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	fmt.Scan(&n)
	ar := make([]data, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i].count, &ar[i].word)
	}
	/* Counting Sort: https://en.wikipedia.org/wiki/Counting_sort#The_algorithm */
	// construct counter array
	SIZE := 100
	counter := make([]int, SIZE)
	for i := 0; i < n; i++ {
		counter[ar[i].count]++
	}
	// sort out indices
	for i := 1; i < SIZE; i++ {
		counter[i] += counter[i-1]
	}
	// produce output array
	output := make([]string, n)
	for i := n - 1; i >= 0; i-- {
		if i >= n/2 {
			output[counter[ar[i].count]-1] = ar[i].word
		} else {
			output[counter[ar[i].count]-1] = "-"
		}
		counter[ar[i].count]--
	}
	fmt.Println(strings.Join(output, " "))
}

type data struct {
	count int
	word  string
}
