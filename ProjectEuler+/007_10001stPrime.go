package main

import "fmt"

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

// refer: 003LargestPrimeFactor.v2.go
// send primes to channel out
func Primes(out chan<- int) {
    out <- 2
    primes := []int{2}
    for n := 3; ; n += 2 {
        ok := true
        for _, p := range primes {
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
