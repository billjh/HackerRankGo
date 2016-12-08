package main

import "fmt"

type seqList struct {
    seq [][]int
}

func main() {
    var n, q int
    fmt.Scanf("%d %d", &n, &q)
    // create seqList and lastAns
    s := seqList{make([][]int, n)}
    for i := 0; i < n; i++ {
        s.seq[i] = []int{}
    }
    var lastAns int
    // read and execute query
    for i := 0; i < q; i++ {
        var a, x, y int
        fmt.Scanf("%d %d %d", &a, &x, &y)
        if a == 1 {
            index := (x ^ lastAns) % n
            s.seq[index] = append(s.seq[index], y)
        } else if a == 2 {
            index := (x ^ lastAns) % n
            lastAns = s.seq[index][y%len(s.seq[index])]
            fmt.Println(lastAns)
        }
    }
}
