package main

import (
    "fmt"
    "math"
)

func main() {
    var plain string
    fmt.Scan(&plain)
    encypt(plain)
}

func encypt(plain string) {
    l := len(plain)
    _, columns := grid(l)
    for i := 0; i < columns; i++ {
        var s string
        for j := i; j < l; j += columns {
            s += string(plain[j])
        }
        fmt.Print(s + " ")
    }
}

func grid(l int) (r, c int) {
    lower := int(math.Floor(math.Sqrt(float64(l))))
    upper := int(math.Ceil(math.Sqrt(float64(l))))
    if lower == upper {
        return lower, lower
    }
    return lower, upper
}
