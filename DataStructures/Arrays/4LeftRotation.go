package main

import "fmt"

// timeout on test case #8
func main() {
    var n, d int
    fmt.Scanf("%d %d", &n, &d)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }
    for i := 0; i < d; i++ {
        arr = leftRotate(arr)
    }
    for i := 0; i < n; i++ {
        fmt.Printf("%d ", arr[i])
    }
}

func leftRotate(arr []int) []int {
    first := arr[0]
    for i := 0; i < len(arr)-1; i++ {
        arr[i] = arr[i+1]
    }
    arr[len(arr)-1] = first
    return arr
}
