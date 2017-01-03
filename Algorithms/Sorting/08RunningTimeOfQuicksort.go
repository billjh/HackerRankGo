package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	fmt.Scan(&n)
	ar := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}
	ar2 := make([]int, n)
	copy(ar2, ar)
	fmt.Println(insertionSort(ar) - quickSortInPlace(ar2, 0, n-1))

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

func quickSortInPlace(ar []int, lo, hi int) int {
	shifts := 0
	if lo < hi {
		p, tmp := partitionInPlace(ar, lo, hi)
		shifts += tmp
		shifts += quickSortInPlace(ar, lo, p-1)
		shifts += quickSortInPlace(ar, p+1, hi)
	}
	return shifts
}

// Lomuto partition scheme
func partitionInPlace(ar []int, lo, hi int) (int, int) {
	shifts := 0
	pivot := ar[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if ar[j] <= pivot {
			ar[i], ar[j] = ar[j], ar[i]
			i++
			shifts++
		}
	}
	ar[i], ar[hi] = ar[hi], ar[i]
	return i, shifts + 1
}
