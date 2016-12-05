package main

import "fmt"

// pass all
func main() {
    var n int
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        var a int
        fmt.Scan(&a)
        fmt.Println(calculate(a))
    }
}

func calculate(n int) int {
    product := 1
    for i := 2; i <= n; i++ {
        product = lcm(i, product)
    }
    return product
}

func lcm(a, b int) int {
    return a * b / gcd(a, b)
}

func gcd(a, b int) int {
    for b > 0 {
        a, b = b, a%b
    }
    return a
}
