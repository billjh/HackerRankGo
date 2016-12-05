package main

import "fmt"

// pass test case #0 #1 #2 #4
// timeout on test case #3 #5
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
