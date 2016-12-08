package main

import (
    "fmt"
    "strings"
)

func main() {
    var n int
    fmt.Scan(&n)
    // use hashmap instead of implementing sparse array
    mp := make(map[string]int)
    for i := 0; i < n; i++ {
        var s string
        fmt.Scanf("%s\n", &s)
        s = strings.TrimSpace(s)
        mp[s] += 1
    }
    var q int
    fmt.Scan(&q)
    for i := 0; i < q; i++ {
        var s string
        fmt.Scanf("%s\n", &s)
        s = strings.TrimSpace(s)
        fmt.Println(mp[s])
    }
}
