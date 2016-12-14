package main

import "fmt"

func main() {
    var n, k int
    fmt.Scan(&n, &k)
    a := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&a[i])
    }
    var charged int
    fmt.Scan(&charged)
    BonAppetit(n, k, a, charged)
}

func BonAppetit(n, k int, a []int, charged int) {
    sum := 0
    for i := 0; i < n; i++ {
        sum += a[i]
    }
    actual := (sum - a[k]) / 2
    if charged == actual {
        fmt.Println("Bon Appetit")
    } else {
        fmt.Println(charged - actual)
    }
}
