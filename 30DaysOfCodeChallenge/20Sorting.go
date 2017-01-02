package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	swaps, first, last := bubbleSort(arr)
	fmt.Println("Array is sorted in", swaps, "swaps.")
	fmt.Println("First Element:", first)
	fmt.Println("Last Element:", last)
}

func bubbleSort(arr []int) (int, int, int) {
	var total int
	for i := 0; i < len(arr); i++ {
		var swaps int
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swaps++
			}
		}
		if swaps == 0 {
			break
		}
		total += swaps
	}
	return total, arr[0], arr[len(arr)-1]
}
