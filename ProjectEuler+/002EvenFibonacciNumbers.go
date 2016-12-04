package main

import "fmt"

// pass all
func main() {
    var n int
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        var a int
        fmt.Scan(&a)
        fmt.Println(sum(a))
    }
}

func sum(limit int) int {
    a, b, sum := 1, 2, 2
    for {
        a, b = b, a+b
        if b >= limit {
            break
        }
        if b%2 == 0 {
            sum += b
        }
    }
    return sum
}
