package main

import "fmt"

func main() {
    var n, k int
    fmt.Scan(&n, &k)
    c := make([]bool, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&c[i])
    }
    e := 100
    for i := k; i <= n; i += k {
        if c[i%n] {
            e -= 3
        } else {
            e -= 1
        }
    }
    fmt.Println(e)
}
