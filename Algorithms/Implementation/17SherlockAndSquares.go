package main

import (
    "fmt"
    "math"
)

func main() {
    var t int
    fmt.Scan(&t)
    for i := 0; i < t; i++ {
        var a, b int
        fmt.Scan(&a, &b)
        fmt.Println(countSquares(a, b))
    }
}

func countSquares(a, b int) int {
    var c int
    for i := int(math.Ceil(math.Sqrt(float64(a)))); i*i <= b; i++ {
        c++
    }
    return c
}
