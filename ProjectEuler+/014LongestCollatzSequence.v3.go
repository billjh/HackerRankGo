package main

import (
    "fmt"
    "math"
)

// pass test case #0 #1 #2 #3 #4
// timeout on test case #5 #6 #7 #8 #9 #10 #11 #12
func main() {
    var t int
    fmt.Scan(&t)
    top, limits := scanInts(t)
    cache := make(map[int]int)
    cache[1] = 1
    chain_history := [][]int{}
    for i := top; i > 1; i-- {
        cur := i
        // build a chain
        chain := []int{}
        for {
            _, ok := cache[cur]
            if ok {
                break
            } else {
                chain = append(chain, cur)
                cur = collatz(cur)
            }
        }
        // cache the chain
        l := len(chain)
        for j := 0; j < l; j++ {
            cache[chain[j]] = cache[cur] + l - j
        }
        if l > 0 {
            chain_history = append(chain_history, chain)
        }
    }
    // get results
    for i := 0; i < t; i++ {
        var cur, length int
        for _, chain := range chain_history {
            for _, node := range chain {
                if length > cache[node] {
                    break
                }
                if node <= limits[i] && cache[node] > cache[cur] {
                    cur, length = node, cache[node]
                    break
                }
            }
        }
        fmt.Println(cur)
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
