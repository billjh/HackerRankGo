package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		if isPrime(n) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not prime")
		}
	}
}

func isPrime(n int) bool {
	if n == 1 { // 1 : false
		return false
	}
	if n <= 3 { // 2, 3 : true
		return true
	}
	if n%2 == 0 { // 4, 6, 8, ... : false
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
