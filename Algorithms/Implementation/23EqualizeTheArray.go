package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }
    fmt.Println(equalize(arr))
}

func equalize(arr []int) int {
    mp := make(map[int]int)
    for _, a := range arr {
        mp[a] += 1
    }
    var _, count int
    for _, v := range mp {
        if v > count {
            count = v
        }
    }
    return len(arr) - count
}
