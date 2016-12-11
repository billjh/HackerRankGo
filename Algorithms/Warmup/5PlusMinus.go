package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }
    plusMinus(arr, n)
}

func plusMinus(arr []int, n int) {
    var pos, neg int
    for i := 0; i < n; i++ {
        if arr[i] > 0 {
            pos++
        } else if arr[i] < 0 {
            neg++
        }
    }
    fmt.Printf("%.6f\n", float32(pos)/float32(n))
    fmt.Printf("%.6f\n", float32(neg)/float32(n))
    fmt.Printf("%.6f\n", float32(n-pos-neg)/float32(n))
}
