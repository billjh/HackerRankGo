package main

import (
    "fmt"
    "math"
)

// pass all
func main() {
    var t int
    fmt.Scan(&t)
    limits := []int{}
    top := 0
    for i := 0; i < t; i++ {
        var n int
        fmt.Scan(&n)
        limits = append(limits, n)
        if n > top {
            top = n
        }
    }
    primes := getPrimes(top)
    for i := 0; i < t; i++ {
        sum := 0
        for j := 0; j < len(primes); j++ {
            p := primes[j]
            if p <= limits[i] {
                sum += p
            } else {
                break
            }
        }
        fmt.Println(sum)
    }
}

func getPrimes(limit int) []int {
    primes := []int{}
    ch := make(chan int)
    go Primes(ch)
    for {
        p := <-ch
        if p <= limit {
            primes = append(primes, p)
        } else {
            break
        }
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
