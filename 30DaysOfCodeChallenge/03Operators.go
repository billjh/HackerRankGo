package main

import (
    "fmt"
    "math"
)

func main() {
    var mealCost float64
    var tipPercent uint32
    var taxPercent uint32

    fmt.Scanf("%g", &mealCost)
    fmt.Scanf("%d", &tipPercent)
    fmt.Scanf("%d", &taxPercent)

    bill := round(mealCost * float64(100+tipPercent+taxPercent) / 100.0)
    fmt.Println("The total meal cost is", bill, "dollars.")
}

func round(f float64) uint32 {
    return uint32(math.Floor(f + .5))
}
