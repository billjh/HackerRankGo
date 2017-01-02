package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    fmt.Println(factorial(n))
}

func factorial(x int) int {
    if x <= 2 {
        return x
    }
    return x * factorial(x-1)
}
