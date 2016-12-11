package main

import "fmt"

// same as 1SimpleArraySum.go
func main() {
    var n int
    fmt.Scan(&n)
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }
    fmt.Println(sum(arr))
}

func sum(arr []int) int {
    var result int
    for _, a := range arr {
        result += a
    }
    return result
}
