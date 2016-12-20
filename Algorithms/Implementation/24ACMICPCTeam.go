package main

import "fmt"

func main() {
    var n, m int
    fmt.Scan(&n, &m)
    teams := make([]string, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&teams[i])
    }
    tp, tm := maxTopics(n, m, teams)
    fmt.Println(tp)
    fmt.Println(tm)
}

func maxTopics(n, m int, teams []string) (int, int) {
    var top, count int
    for i := 0; i < n-1; i++ {
        for j := i + 1; j < n; j++ {
            ts := teamScore(teams[i], teams[j], m)
            if ts > top {
                top = ts
                count = 1
            } else if ts == top {
                count++
            }
        }
    }
    return top, count
}

func teamScore(a, b string, m int) int {
    s := 0
    for i := 0; i < m; i++ {
        if a[i] == '1' || b[i] == '1' {
            s++
        }
    }
    return s
}
