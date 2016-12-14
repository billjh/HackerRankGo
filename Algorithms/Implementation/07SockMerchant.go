package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    c := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&c[i])
    }
    fmt.Println(sellSocks(c))
}

func sellSocks(c []int) int {
    cm := make(map[int]int)
    for _, color := range c {
        cm[color] += 1
    }
    pairs := 0
    for _, number := range cm {
        pairs += number / 2
    }
    return pairs
}
