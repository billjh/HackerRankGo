package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    rd := bufio.NewReader(os.Stdin)
    li, _ := rd.ReadString('\n')
    fmt.Println("Hello, World.")
    fmt.Print(li)
}
