package main

import (
    "fmt"
    "math"
)

func main() {
    var n, m int
    fmt.Scanf("%d %d", &n, &m)
    a_max, _, a := scanInts(n)
    _, b_min, b := scanInts(m)
    count := 0
    for i := a_max; i <= b_min; i += a_max {
        if between(i, a, b) {
            count++
        }
    }
    fmt.Println(count)
}

func between(i int, a, b []int) bool {
    for _, x := range a {
        if i%x != 0 {
            return false
        }
    }
    for _, y := range b {
        if y%i != 0 {
            return false
        }
    }
    return true
}

func scanInts(n int) (top int, bottom int, arr []int) {
    arr = make([]int, n)
    top = math.MinInt64
    bottom = math.MaxInt64
    for i := 0; i < n; i++ {
        var m int
        fmt.Scan(&m)
        arr[i] = m
        if m > top {
            top = m
        }
        if m < bottom {
            bottom = m
        }
    }
    return
}
