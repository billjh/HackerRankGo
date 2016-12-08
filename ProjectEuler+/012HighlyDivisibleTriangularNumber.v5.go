package main

import (
    "fmt"
    "math"
)

// pass all
func main() {
    var t int
    fmt.Scan(&t)
    _, arr := scanInts(t)
    results := make([]int, t)
    var done int
    for i := 1; done != t; i++ {
        // i and i+1 are co-prime
        var next int
        if i%2 == 0 {
            next = countFactors(i/2) * countFactors(i+1)
        } else {
            next = countFactors(i) * countFactors((i+1)/2)
        }
        for j := 0; j < t; j++ {
            if results[j] == 0 && next > arr[j] {
                results[j] = i * (i + 1) / 2
                done++
            }
        }
    }
    for _, r := range results {
        fmt.Println(r)
    }
}

func countFactors(n int) int {
    total, num := 1, n
    // fake prime will do
    // 2, 3, 5, 7, 9, ...
    for num%2 == 0 {
        total += 1
        num /= 2
    }
    for p := 3; num > 1; p += 2 {
        count := 1
        for num%p == 0 {
            count += 1
            num /= p
        }
        total *= count
    }
    return total
}

// refer to: 012HighlyDivisibleTriangularNumber.go
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
