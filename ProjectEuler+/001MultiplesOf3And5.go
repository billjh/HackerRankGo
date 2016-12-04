package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        var a uint
        fmt.Scan(&a)
        fmt.Println(calculate(a))
    }
}

func calculate(limit uint) uint {
    return sum(limit, 3) + sum(limit, 5) - sum(limit, 3*5)
}

func sum(limit uint, base uint) uint {
    count := (limit - 1) / base
    return count * base * (count + 1) / 2
}
