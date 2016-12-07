package main

import "fmt"
import "math/big"

// pass all
func main() {
    var n int
    fmt.Scan(&n)
    sum, a := big.NewInt(0), big.NewInt(0)
    var s string
    for i := 0; i < n; i++ {
        fmt.Scanf("%s\n", &s)
        a.SetString(s, 10)
        sum.Add(sum, a)
    }
    fmt.Println(sum.String()[:10])
}
