package main

import "fmt"

func main() {
    var s string
    var n int
    fmt.Scan(&s, &n)
    fmt.Println(countA(s, n))
}

func countA(s string, n int) int {
    var all, tail int
    for i := 0; i < len(s); i++ {
        if s[i] != 'a' {
            continue
        }
        all++
        if i < n%len(s) {
            tail++
        }
    }
    return all*(n/len(s)) + tail
}
