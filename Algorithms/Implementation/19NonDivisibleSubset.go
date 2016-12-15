package main

import "fmt"

func main() {
    var n, k int
    fmt.Scan(&n, &k)
    s := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&s[i])
    }
    fmt.Println(subsetSize(s, k))
}

func subsetSize(s []int, k int) int {
    mp := make(map[int]int)
    for _, e := range s {
        mp[e%k] += 1
    }
    size := 0
    if mp[0] > 0 {
        size += 1
    }
    for i := 1; float64(i) < float64(k)/2.0; i++ {
        if mp[i] > mp[k-i] {
            size += mp[i]
        } else {
            size += mp[k-i]
        }
    }
    if k%2 == 0 && mp[k/2] > 0 {
        size += 1
    }
    return size
}
