package main

import "fmt"

func main() {
    var hh, mm, ss int
    var apm string
    fmt.Scanf("%d:%d:%d%s", &hh, &mm, &ss, &apm)
    if apm == "PM" && hh != 12 || apm == "AM" && hh == 12 {
        hh = (hh + 12) % 24
    }
    fmt.Printf("%02d:%02d:%02d", hh, mm, ss)
}
