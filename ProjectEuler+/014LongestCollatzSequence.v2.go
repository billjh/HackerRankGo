package main

import (
    "fmt"
    "math"
)

var ENTRY_LIMIT int = 100000000

// pass test case #0 #1 #2 #3 #4
// timeout on test case #5 #6 #7 #8 #9 #10 #11 #12
func main() {
    var t int
    fmt.Scan(&t)
    top, limits := scanInts(t)
    // use array instead of map to cache
    cache := make([]int, ENTRY_LIMIT)
    cache[1] = 1
    results := make([]int, t)
    for i := top; i > 1; i-- {
        cur := i
        // build a chain
        chain := []int{}
        for {
            if cur < ENTRY_LIMIT && cache[cur] != 0 {
                break
            } else {
                chain = append(chain, cur)
                cur = collatz(cur)
            }
        }
        // count the chain
        l := len(chain)
        for j := 0; j < l; j++ {
            if chain[j] < ENTRY_LIMIT {
                cache[chain[j]] = cache[cur] + l - j
            }
            // write to results array
            for k := 0; k < t; k++ {
                c := chain[j]
                if c <= limits[k] && cache[c] > cache[results[k]] {
                    results[k] = c
                }
            }
        }
    }
    for i := 0; i < t; i++ {
        fmt.Println(results[i])
    }
}

func collatz(n int) int {
    if n%2 == 0 {
        return n / 2
    } else {
        return 3*n + 1
    }
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
