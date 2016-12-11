package main

import "fmt"

func main() {
    var n, k, q int
    fmt.Scan(&n, &k, &q)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }
    arr = rotate(arr, k)
    for i := 0; i < q; i++ {
        var query int
        fmt.Scan(&query)
        fmt.Println(arr[query])
    }
}

func rotate(arr []int, k int) []int {
    l := len(arr) - k%len(arr)
    return append(append([]int{}, arr[l:]...), arr[:l]...)
}
