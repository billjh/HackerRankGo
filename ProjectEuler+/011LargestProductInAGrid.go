package main

import "fmt"

const GRID_SIZE = 20

// pass all
func main() {
    grid := make([][]int, GRID_SIZE)
    for i := 0; i < GRID_SIZE; i++ {
        grid[i] = make([]int, GRID_SIZE)
        for j := 0; j < GRID_SIZE; j++ {
            fmt.Scan(&grid[i][j])
        }
    }
    fmt.Println(calculate(grid, GRID_SIZE, 4))
}

func calculate(grid [][]int, n, k int) int {
    max := 0
    // left <-> right
    for i := 0; i < n; i++ {
        for j := 0; j < n-k+1; j++ {
            product := 1
            for l := 0; l < k && product > 0; l++ {
                product *= grid[i][j+l]
            }
            if max < product {
                max = product
            }
        }
    }
    // up <-> down
    for i := 0; i < n-k+1; i++ {
        for j := 0; j < n; j++ {
            product := 1
            for l := 0; l < k && product > 0; l++ {
                product *= grid[i+l][j]
            }
            if max < product {
                max = product
            }
        }
    }
    // left-up <-> right-down
    for i := 0; i < n-k+1; i++ {
        for j := 0; j < n-k+1; j++ {
            product := 1
            for l := 0; l < k && product > 0; l++ {
                product *= grid[i+l][j+l]
            }
            if max < product {
                max = product
            }
        }
    }
    // left-down <-> right-up
    for i := 0; i < n-k+1; i++ {
        for j := k - 1; j < n; j++ {
            product := 1
            for l := 0; l < k && product > 0; l++ {
                product *= grid[i+l][j-l]
            }
            if max < product {
                max = product
            }
        }
    }
    return max
}
