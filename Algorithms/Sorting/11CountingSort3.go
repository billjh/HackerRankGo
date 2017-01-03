package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	fmt.Scan(&n)
	SIZE := 100
	results := make([]int, SIZE)
	for i := 0; i < n; i++ {
		var c int
		var s string
		fmt.Scan(&c, &s)
		results[c]++
	}
	sum := 0
	for i := 0; i < SIZE; i++ {
		sum += results[i]
		fmt.Printf("%d ", sum)
	}
}
