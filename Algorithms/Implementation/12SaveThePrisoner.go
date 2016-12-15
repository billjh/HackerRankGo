package main

import "fmt"

func main() {
    var t int
    fmt.Scan(&t)
    for i := 0; i < t; i++ {
        var n, m, s int
        fmt.Scan(&n, &m, &s)
        fmt.Println((s+m-2)%n + 1)
    }
}
