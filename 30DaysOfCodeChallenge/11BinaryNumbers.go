package main

import (
    "fmt"
    "strconv"
)

func main() {
    var n int64
    fmt.Scan(&n)
    bs := strconv.FormatInt(n, 2)
    var final, counter int
    for i := 0; i < len(bs); i++ {
        b := bs[i]
        if b == '1' {
            counter++
            final = max(final, counter)
        } else {
            counter = 0
        }
    }
    fmt.Println(final)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
