package main

import (
    "testing"
)

func BenchmarkPart1(b *testing.B) {
    cMap, maxCoord, _:= parse("input2.txt")
    for i := 0; i < b.N; i++ {
        sol(cMap, maxCoord, 1)
    }
}

func BenchmarkPart2(b *testing.B) {
    cMap, maxCoord, _:= parse("input2.txt")
    for i := 0; i < b.N; i++ {
        sol(cMap, maxCoord, 2)
    }
}
