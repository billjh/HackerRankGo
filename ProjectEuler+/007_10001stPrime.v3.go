package main

import (
    "fmt"
    "math"
)

// pass all
func main() {
    var t int
    fmt.Scan(&t)
    top, ranks := scanInts(t)
    // generate primes to an array
    primes := genPrimes(top)
    for i := 0; i < t; i++ {
        fmt.Println(primes[ranks[i]-1])
    }
}

func genPrimes(limit int) []int {
    primes := []int{}
    ch := make(chan int)
    go Primes(ch)
    for i := 0; i < limit; i++ {
        primes = append(primes, <-ch)
    }
    return primes
}

// refer: 007_10001stPrime.v2.go
// send primes to channel out
func Primes(out chan<- int) {
    out <- 2
    out <- 3
    primes := []int{2, 3}
    for n := 5; ; n += 2 {
        ok := true
        for _, p := range primes {
            // speed-up:
            // if the divisor become larger than square root of the candidate
            // then the candidate is confirmed to be a prime
            if float64(p) > math.Sqrt(float64(n)) {
                break
            }
            if n%p == 0 {
                ok = false
                break
            }
        }
        if ok {
            primes = append(primes, n)
            out <- n
        } else {
            continue
        }
    }
}

// refactored with: 012HighlyDivisibleTriangularNumber.go
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
