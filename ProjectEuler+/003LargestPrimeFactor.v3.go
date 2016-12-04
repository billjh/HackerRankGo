package main

import (
    "fmt"
    "math"
)

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
    result := 2
    for cur%2 == 0 {
        cur /= 2
    }
    for fakeprime := 3; ; fakeprime += 2 {
        if fakeprime > cur {
            return result
        }
        if float64(fakeprime) > math.Sqrt(float64(n)) {
            return cur
        }
        for cur%fakeprime == 0 {
            result = fakeprime
            cur /= fakeprime
        }
    }
}
