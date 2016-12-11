package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    mat := make([][]int, n)
    for i := 0; i < n; i++ {
        mat[i] = make([]int, n)
        for j := 0; j < n; j++ {
            fmt.Scan(&mat[i][j])
        }
    }
    fmt.Println(diagonalDifference(mat, n))
}

func diagonalDifference(mat [][]int, n int) int {
    var a, b int
    for i := 0; i < n; i++ {
        a += mat[i][i]
        b += mat[n-i-1][i]
    }
    if a > b {
        return a - b
    }
    return b - a
}
