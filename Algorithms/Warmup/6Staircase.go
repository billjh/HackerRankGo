package main

import (
    "fmt"
    "strings"
)

func main() {
    var n int
    fmt.Scan(&n)
    staircase(n)
}

func staircase(n int) {
    for i := 1; i <= n; i++ {
        fmt.Println(strings.Repeat(" ", n-i) + strings.Repeat("#", i))
    }
}
