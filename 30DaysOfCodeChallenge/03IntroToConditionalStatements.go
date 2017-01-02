package main

import "fmt"

func main() {
    var n uint8
    fmt.Scanf("%d", &n)
    fmt.Println(check(n))
}

func check(n uint8) string {
    if n < 1 || n > 100 {
        panic(n)
    }
    if n%2 == 1 || (n >= 6 && n <= 20) {
        return "Weird"
    }
    return "Not Weird"
}
