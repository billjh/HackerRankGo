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
    for _, n := range arr {
        fmt.Println(findFirst(n))
    }
}

// find the first triangle number has more factors than n
func findFirst(n int) int {
    for i := 1; ; i++ {
        factors := union(factorize(i), factorize(i+1))
        factors[2] -= 1
        if count(factors) > n {
            return i * (i + 1) / 2
        }
    }
}

// union two maps
func union(a, b map[int]int) map[int]int {
    for key, value := range b {
        _, ok := a[key]
        if ok {
            a[key] += value
        } else {
            a[key] = value
        }
    }
    return a
}

func count(factors map[int]int) int {
    count := 1
    for _, times := range factors {
        count *= (1 + times)
    }
    return count
}

// factorize and return prime:count as a map
func factorize(n int) map[int]int {
    cur := n
    factors := make(map[int]int)
    ch := make(chan int)
    go Primes(ch)
    for cur > 1 {
        p := <-ch
        if cur%p == 0 {
            cur /= p
            factors[p] = 1
        }
        for cur%p == 0 {
            cur /= p
            factors[p] += 1
        }
    }
    return factors
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
