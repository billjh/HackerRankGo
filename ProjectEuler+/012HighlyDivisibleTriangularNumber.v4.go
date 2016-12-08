package main

import (
    "fmt"
    "math"
)

// pass test case #0 #1 #2 #3
// timeout on test case #4 #5 #6 #7
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
    ch := make(chan int)
    go Primes(ch)
    for p := <-ch; num > 1; p = <-ch {
        count := 1
        for num%p == 0 {
            count += 1
            num /= p
        }
        total *= count
    }
    return total
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
