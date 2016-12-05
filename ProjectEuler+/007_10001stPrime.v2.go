package main

import (
    "fmt"
    "math"
)

// pass test case #0 #1 #2
// timeout on test case #3 #4
func main() {
    var t int
    fmt.Scan(&t)
    for i := 0; i < t; i++ {
        var n int
        fmt.Scan(&n)
        ch := make(chan int)
        go Primes(ch)
        for j := 0; j < n-1; j++ {
            <-ch
        }
        fmt.Println(<-ch)
    }
}

// enhanced version over 003LargestPrimeFactor.v2.go
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
