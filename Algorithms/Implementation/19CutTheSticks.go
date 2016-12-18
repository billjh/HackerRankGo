package main

import (
    "fmt"
    "sort"
)

func main() {
    var n int
    fmt.Scan(&n)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }
    cut(arr)
}

func cut(arr []int) {
    sort.Ints(arr)
    for i := len(arr) - 1; i > 0; i-- {
        arr[i] -= arr[i-1]
    }
    for i := 0; i < len(arr); i++ {
        if arr[i] == 0 {
            continue
        }
        fmt.Println(len(arr) - i)
    }
}
