package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	// number of test cases
	fmt.Println(5)
	// case 1: expect YES
	fmt.Println("5 4")
	fmt.Println("-2 -1 0 1 2")
	// case 2: expect NO
	fmt.Println("5 3")
	fmt.Println("-2 -1 0 1 2")
	// case 3: expect YES
	fmt.Println("5 4")
	fmt.Println("2 1 0 -1 -2")
	// case 4: expect NO
	fmt.Println("5 1")
	fmt.Println("2 1 0 -1 -2")
	// case 5: expect YES
	fmt.Println("5 3")
	fmt.Println("-1 0 1 2 3")
}
