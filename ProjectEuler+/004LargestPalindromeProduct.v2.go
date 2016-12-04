package main

import (
    "fmt"
    "strconv"
)

// pass all
func main() {
    var n int
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        var limit, cur int
        fmt.Scan(&limit)
        for i := 100; i < 1000; i++ {
            for j := i + 1; j < 1000; j++ {
                k := i * j
                if k >= limit {
                    break
                }
                // Hint from discussions from hackerrank:
                // all 6-digit palindrome is multiple of 11
                // speed up by checking before doing expensive string operations
                // Proof:
                // 100001 * a + 10010 * b + 1100 * c => 11 * (9091 * a + 910 * b + 100 * c)
                if k%11 == 0 && isPalindrome(k) {
                    cur = max(k, cur)
                }
            }
        }
        fmt.Println(cur)
    }
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func isPalindrome(n int) bool {
    s := strconv.Itoa(n)
    return s == reverse(s)
}

func reverse(s string) string {
    l := len(s)
    t := []rune(s)
    for i := 0; i < l/2; i++ {
        t[i], t[l-1-i] = t[l-1-i], t[i]
    }
    return string(t)
}
