package main

import "fmt"

// pass all
func main() {
    var n, d int
    fmt.Scanf("%d %d", &n, &d)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }
    arr = leftRotateN(arr, d)
    for i := 0; i < n; i++ {
        fmt.Printf("%d ", arr[i])
    }
}

func leftRotateN(arr []int, n int) []int {
    m := n % len(arr)
    return append(append([]int{}, arr[m:]...), arr[:m]...)
}
