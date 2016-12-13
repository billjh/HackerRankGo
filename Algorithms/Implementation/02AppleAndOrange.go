package main

import "fmt"

func main() {
    var s, t, a, b, m, n, apples, oranges int
    fmt.Scanf("%d %d\n%d %d\n%d %d\n", &s, &t, &a, &b, &m, &n)
    for i := 0; i < m; i++ {
        var d int
        fmt.Scan(&d)
        if a+d >= s && a+d <= t {
            apples++
        }
    }
    for i := 0; i < n; i++ {
        var d int
        fmt.Scan(&d)
        if b+d >= s && b+d <= t {
            oranges++
        }
    }
    fmt.Println(apples)
    fmt.Println(oranges)
}
