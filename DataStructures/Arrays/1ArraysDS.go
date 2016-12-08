package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }
    for _, a := range reverse(arr) {
        fmt.Printf("%d ", a)
    }
}

func reverse(arr []int) []int {
    l := len(arr)
    for i := 0; i < l/2; i++ {
        arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
    }
    return arr
}
