package main

import (
    "fmt"
    "strconv"
)

func main() {
    var t int
    fmt.Scan(&t)
    for i := 0; i < t; i++ {
        var n int
        fmt.Scan(&n)
        c := 0
        for _, d := range []rune(strconv.Itoa(n)) {
            g, _ := strconv.Atoi(string(d))
            if g > 0 && n%(g) == 0 {
                c++
            }
        }
        fmt.Println(c)
    }
}
