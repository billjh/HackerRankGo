package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    sum, cur := 0, 2
    for i := 0; i < n; i++ {
        sum += cur
        cur = 3 * cur / 2
    }
    fmt.Println(sum)
}
