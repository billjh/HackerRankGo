package main

import "fmt"

// pass test case #0 #1 #2 #3
// timeout on test case #4 #5
func main() {
    var n int
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        var a int
        fmt.Scan(&a)
        for n := a; n >= 2; n-- {
            if a%n == 0 {
                if isPrime(n) {
                    fmt.Println(n)
                    break
                }
            }
        }
    }
}

func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i < n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}
