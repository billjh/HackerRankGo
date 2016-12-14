package main

import "fmt"

func main() {
    var x1, v1, x2, v2 int
    fmt.Scan(&x1, &v1, &x2, &v2)
    fmt.Println(kangaroo(x1, v1, x2, v2))
}

func kangaroo(x1, v1, x2, v2 int) string {
    if (x1-x2)*(v1-v2) > 0 || (v1-v2) == 0 || (x1-x2)%(v2-v1) != 0 {
        return "NO"
    }
    return "YES"
}
