package main

import (
    "fmt"
    "math"
)

// pass test case #0 #1 #2
// timeout on test case #3 #4 #5 #6 #7
func main() {
    var t int
    fmt.Scan(&t)
    _, arr := scanInts(t)
    for _, n := range arr {
        fmt.Println(findFirst(n))
    }
}

// find the first triangle number has more factors than n
func findFirst(n int) int {
    tri := 0
    for i := 1; ; i++ {
        tri += i
        fac := countFactors(tri)
        if fac > n {
            return tri
        }
    }
}

func countFactors(n int) int {
    count := 0
    for i := 1; i <= n/2; i++ {
        if n%i == 0 {
            count++
        }
    }
    // n is a factor of n
    if n >= 2 {
        count++
    }
    return count
}

// scan a number of integers
// return the maximum and an array
func scanInts(n int) (top int, arr []int) {
    arr = make([]int, n)
    top = math.MinInt64
    for i := 0; i < n; i++ {
        var m int
        fmt.Scan(&m)
        arr[i] = m
        if m > top {
            top = m
        }
    }
    return
}
