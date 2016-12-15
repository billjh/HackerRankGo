package main

import (
    "fmt"
    "math/big"
)

func main() {
    var n int
    fmt.Scan(&n)
    f := big.NewInt(int64(1))
    for i := 2; i <= n; i++ {
        f.Mul(f, big.NewInt(int64(i)))
    }
    fmt.Println(f)
}
