package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    hm := make(map[string]int)
    var num int
    var name string
    for i := 0; i < n; i++ {
        fmt.Scanf("%s %d", &name, &num)
        hm[name] = num
    }
    for {
        _, err := fmt.Scanf("%s", &name)
        if err != nil {
            break
        }
        num, ok := hm[name]
        if ok {
            fmt.Printf("%s=%d\n", name, num)
        } else {
            fmt.Println("Not found")
        }
    }
}
