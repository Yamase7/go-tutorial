package main

import"fmt"

func main() {
    // int64 用 map の初期化
    ints := map[string]int64{
        "first":  34,
        "second": 12,
    }

    // float 用 map の初期化
    floats := map[string]float64{
        "first":  35.98,
        "second": 26.99,
    }

    fmt.Printf("Non-Generic Sums: %v and %v\n",
        SumInts(ints),
        SumFloats(floats))
}

// mの値を加算していく
func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

// mの値を加算していく
func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}