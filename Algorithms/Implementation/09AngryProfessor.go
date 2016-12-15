package main

import "fmt"

func main() {
    var t int
    fmt.Scan(&t)
    for i := 0; i < t; i++ {
        var n, k int
        fmt.Scan(&n, &k)
        for j := 0; j < n; j++ {
            var a int
            fmt.Scan(&a)
            if a <= 0 {
                k--
            }
        }
        if k <= 0 {
            fmt.Println("NO")
        } else {
            fmt.Println("YES")
        }
    }
}
