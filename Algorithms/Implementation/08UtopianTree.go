package main

import "fmt"

func main() {
    var t int
    fmt.Scan(&t)
    for i := 0; i < t; i++ {
        var n int
        fmt.Scan(&n)
        fmt.Println(utopianTree(n))
    }
}

func utopianTree(n int) int {
    h := 1
    for i := 0; i < n; i++ {
        if i%2 == 0 {
            h *= 2
        } else {
            h++
        }
    }
    return h
}
