package main

import (
    "fmt"
    "strconv"
)

// pass all
func main() {
    var t int
    fmt.Scan(&t)
    for i := 0; i < t; i++ {
        var n, k int
        var digits string
        fmt.Scanf("%d %d\n%s\n", &n, &k, &digits)
        fmt.Println(largestProduct(digits, n, k))
    }
}

func largestProduct(digits string, n, k int) int {
    max := 0
    for i := 0; i < n-k+1; i++ {
        p := calculate(digits[i : i+k])
        if p > max {
            max = p
        }
    }
    return max
}

func calculate(digits string) int {
    product := 1
    for _, d := range digits {
        i, _ := strconv.Atoi(string(d))
        product *= i
    }
    return product
}
