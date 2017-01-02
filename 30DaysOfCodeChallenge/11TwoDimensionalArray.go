package main

import (
    "fmt"
    "math"
)

const ARRAY_SIZE int = 6

func main() {
    arr := make([][]int, ARRAY_SIZE)
    for i := 0; i < ARRAY_SIZE; i++ {
        arr[i] = make([]int, ARRAY_SIZE)
        for j := 0; j < ARRAY_SIZE; j++ {
            fmt.Scanf("%d", &arr[i][j])
        }
    }
    result := math.MinInt64
    for i := 0; i < ARRAY_SIZE-2; i++ {
        for j := 0; j < ARRAY_SIZE-2; j++ {
            // hour glass sum
            sum := arr[i][j] + arr[i][j+1] + arr[i][j+2] +
                arr[i+1][j+1] +
                arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
            if sum > result {
                result = sum
            }
        }
    }
    fmt.Println(result)
}
