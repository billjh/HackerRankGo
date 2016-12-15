package main

import "fmt"

func main() {
    var s, t string
    var k int
    fmt.Scan(&s, &t, &k)
    sl, tl, cl := len([]rune(s)), len([]rune(t)), common(s, t)
    if appendAndDelete(sl, tl, cl, k) {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}

func appendAndDelete(sl, tl, cl, k int) bool {
    if k >= sl+tl {
        return true
    }
    if k >= sl+tl-cl*2 && (k-(sl+tl-cl*2))%2 == 0 {
        return true
    }
    return false
}

func common(s, t string) int {
    sr, rr := []rune(s), []rune(t)
    var l, c int
    if len(sr) > len(rr) {
        l = len(rr)
    } else {
        l = len(sr)
    }
    for i := 0; i < l; i++ {
        if sr[i] == rr[i] {
            c++
        } else {
            break
        }
    }
    return c
}
