package main

import (
	"fmt"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	SIZE := 100
	var n int
	fmt.Scan(&n)
	counter := make([]int, SIZE)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		counter[a]++
	}
	fmt.Println(strings.Trim(fmt.Sprint(counter), "[]"))
}
