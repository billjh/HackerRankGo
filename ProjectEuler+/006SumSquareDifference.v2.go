package main

import "fmt"

// pass all
func main() {
    var t int
    fmt.Scan(&t)
    for i := 0; i < t; i++ {
        var n int
        fmt.Scan(&n)
        sum := 0
        for j := 1; j <= n-1; j++ {
            // optimize sum from 2 * j * (j + 1) to 2 * j * n
            sum += j * (j + 1 + n) * (n - j)
        }
        fmt.Println(sum)
    }
}
