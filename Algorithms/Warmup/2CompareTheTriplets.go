package main

import "fmt"

func main() {
    var a0, a1, a2, b0, b1, b2, alice, bob int
    fmt.Scanf("%d %d %d\n", &a0, &a1, &a2)
    fmt.Scanf("%d %d %d\n", &b0, &b1, &b2)
    if a0 > b0 {
        alice++
    } else if a0 < b0 {
        bob++
    }
    if a1 > b1 {
        alice++
    } else if a1 < b1 {
        bob++
    }
    if a2 > b2 {
        alice++
    } else if a2 < b2 {
        bob++
    }
    fmt.Println(alice, bob)
}
