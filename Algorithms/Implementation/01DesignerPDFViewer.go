package main

import "fmt"

func main() {
    var SIZE int = 26
    heights := make([]int, SIZE)
    for i := 0; i < SIZE; i++ {
        fmt.Scan(&heights[i])
    }
    var s string
    fmt.Scanf("%s\n", &s)
    var maxh int
    for _, c := range s {
        h := heights[(int(c)-int('a'))%26]
        if h > maxh {
            maxh = h
        }
    }
    fmt.Println(maxh * len(s))
}
