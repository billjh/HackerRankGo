package main

import "fmt"

// pass test case #0
// timeout on test case #1
func main() {
    var t int
    fmt.Scan(&t)
    for i := 0; i < t; i++ {
        var n int
        fmt.Scan(&n)
        var sum uint = 0
        // (a + b + c + ...)^2 - (a^2 + b^2 + c^2 + ...)
        // = 2 * (ab + ac + ... ) + 2 * (bc + ...) + ...
        for j := 1; j <= n-1; j++ {
            for k := j + 1; k <= n; k++ {
                sum += uint(2 * j * k)
            }
        }
        fmt.Println(sum)
    }
}
