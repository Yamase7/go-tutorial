package main

import"fmt"

type Number interface {
	int64 | float64
}

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
    fmt.Printf("Generic Sums: %v and %v\n",
        SumIntsOrFloats[string, int64](ints),
        SumIntsOrFloats[string, float64](floats))
    fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
        SumIntsOrFloats(ints),
        SumIntsOrFloats(floats))
	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
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

// map の m の値をすべて合計した値を返す。IntとFloatの両方に対応している。
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

// map mの値を合計する。map の値として整数と浮動小数点数の両方をサポートしている。
func SumNumbers[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
