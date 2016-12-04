package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    var n int
    fmt.Scan(&n)
    r := bufio.NewReader(os.Stdin)
    for i := 0; i < n; i++ {
        s, _ := r.ReadString('\n')
        s = strings.TrimSpace(s)
        var oddStr, evenStr string
        for j := 0; j < len(s); j += 2 {
            oddStr += string(s[j])
        }
        if len(s) > 1 {
            for j := 1; j < len(s); j += 2 {
                evenStr += string(s[j])
            }
        }
        fmt.Println(oddStr, evenStr)
    }
}
