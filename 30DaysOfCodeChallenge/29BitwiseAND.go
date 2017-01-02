package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n, k int
		fmt.Scan(&n, &k)
		fmt.Println(calculate(n, k))
	}
}

func calculate(n, k int) int {
	var result int
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			a := i & j
			if a < k && a > result {
				result = a
			}
		}
	}
	return result
}
