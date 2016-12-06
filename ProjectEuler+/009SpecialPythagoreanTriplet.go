package main

import "fmt"

// pass all
func main() {
    var t int
    fmt.Scan(&t)
    for i := 0; i < t; i++ {
        var n int
        fmt.Scan(&n)
        fmt.Println(findLargest(n))
    }
}

func findLargest(n int) int {
    max := -1
    // a >= b >= c
    for a := n / 3; a < n/2; a++ {
        for c := 1; c < (n-a)/2+1; c++ {
            b := n - a - c
            if a*a == b*b+c*c && a*b*c > max {
                max = a * b * c
            }
        }
    }
    return max
}
