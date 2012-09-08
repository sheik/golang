package main

import (
    "fmt"
    "math"
)

func main() {
    // calculate pi using 5000 summations
    fmt.Println(pi(5000))
}

func pi(n int) float64 {

    // calculate n pi terms
    ch := make(chan float64)
    for k := 0; k <= n; k++ {
        go term(ch, float64(k))
    }

    // add terms together
    f := 0.0
    for k:= 0; k <= n; k++ {
        f+= <-ch
    }
    return f
}

// calculate one term for pi
func term(ch chan float64, k float64) {
    ch <- 4 * math.Pow(-1, k) / (2 * k + 1)
}
