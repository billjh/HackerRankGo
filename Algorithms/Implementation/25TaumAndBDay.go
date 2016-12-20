package main

import "fmt"

func main() {
    var t int
    fmt.Scan(&t)
    for i := 0; i < t; i++ {
        var b, w, x, y, z int
        fmt.Scan(&b, &w, &x, &y, &z)
        fmt.Println(buyGifts(b, w, x, y, z))
    }
}

func buyGifts(b, w, x, y, z int) int {
    costs := 0
    if y+z >= x {
        costs += b * x
    } else {
        costs += b * (y + z)
    }
    if x+z >= y {
        costs += w * y
    } else {
        costs += w * (x + z)
    }
    return costs
}
