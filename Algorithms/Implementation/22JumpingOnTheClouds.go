package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    cl := make([]bool, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&cl[i])
    }
    fmt.Println(jumps(cl, n))
}

func jumps(cl []bool, n int) int {
    jps := 0
    i := 0
    for i < n-2 {
        if cl[i+2] == false {
            i += 2
        } else {
            i++
        }
        jps++
    }
    if i == n-2 {
        jps++
    }
    return jps
}
