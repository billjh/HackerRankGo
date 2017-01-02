package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var i uint32 = 4
    var d float32 = 4.0
    var s = "HackerRank "

    scanner := bufio.NewReader(os.Stdin)

    var i2 uint32
    fmt.Scanf("%d", &i2)

    var d2 float32
    fmt.Scanf("%g", &d2)

    s2, _ := scanner.ReadString('\n')

    fmt.Println(i + i2)
    fmt.Printf("%.1f\n", d+d2)
    fmt.Println(s + s2)
}
