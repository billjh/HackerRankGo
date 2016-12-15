package main

import (
    "fmt"
    "strconv"
)

func main() {
    var i, j, k int
    fmt.Scan(&i, &j, &k)
    count := 0
    for d := i; d <= j; d++ {
        if isBeautiful(d, k) {
            count++
        }
    }
    fmt.Println(count)
}

func isBeautiful(d, k int) bool {
    b := reversed(d)
    if d > b {
        return (d-b)%k == 0
    } else {
        return (b-d)%k == 0
    }
}

func reversed(d int) int {
    ds := []rune(strconv.Itoa(d))
    for i := 0; i < len(ds)/2; i++ {
        ds[i], ds[len(ds)-i-1] = ds[len(ds)-i-1], ds[i]
    }
    b, _ := strconv.Atoi(string(ds))
    return b
}
