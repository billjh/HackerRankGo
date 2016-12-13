package main

import (
    "fmt"
    "math"
)

func main() {
    var min, max, all int
    min = math.MaxInt64
    for {
        var n int
        _, err := fmt.Scan(&n)
        if err != nil {
            break
        }
        if n < min {
            min = n
        }
        if n > max {
            max = n
        }
        all += n
    }
    if min == math.MaxInt64 && max == 0 { // no input
        fmt.Println(0, 0)
    } else {
        fmt.Println(all-max, all-min)
    }
}
