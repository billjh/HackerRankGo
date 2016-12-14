package main

import "fmt"

func main() {
    var n, k int
    fmt.Scanf("%d %d", &n, &k)
    a := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&a[i])
    }
    fmt.Println(countPairs(k, a))
}

func countPairs(k int, a []int) int {
    count := 0
    for i := 0; i < len(a)-1; i++ {
        for j := i + 1; j < len(a); j++ {
            if (a[i]+a[j])%k == 0 {
                count++
            }
        }
    }
    return count
}
