package main

import "fmt"

func main() {
	var aD, aM, aY, eD, eM, eY int
	fmt.Scan(&aD, &aM, &aY, &eD, &eM, &eY)
	fmt.Println(fine(aD, aM, aY, eD, eM, eY))
}

func fine(aD, aM, aY, eD, eM, eY int) int {
	if aY > eY {
		return 10000
	}
	if aY == eY && aM > eM {
		return 500 * (aM - eM)
	}
	if aY == eY && aM == eM && aD > eD {
		return 15 * (aD - eD)
	}
	return 0
}
