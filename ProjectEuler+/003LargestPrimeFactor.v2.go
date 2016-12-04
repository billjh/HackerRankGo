package main

import (
    "fmt"
    "math"
)

// timeout on test case #5
func main() {
    var n int
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        var a int
        fmt.Scan(&a)
        fmt.Println(largestPrimeFactor(a))
    }
}

func largestPrimeFactor(n int) int {
    cur := n
    result := 0
    ch := make(chan int)
    go Primes(ch)
    for {
        prime := <-ch
        if prime > cur {
            return result
        }
        for cur%prime == 0 {
            result = prime
            cur /= prime
        }
    }
}

// send primes to channel out
func Primes(out chan<- int) {
    out <- 2
    primes := []int{2}
    for n := 3; ; n += 2 {
        for _, p := range primes {
            if float64(p) > math.Sqrt(float64(n)) {
                break
            }
            if n%p == 0 {
                continue
            }
        }
        primes = append(primes, n)
        out <- n
    }
}
