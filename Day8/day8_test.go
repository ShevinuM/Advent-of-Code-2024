package main

import (
    "testing"
)

func BenchmarkPart1(b *testing.B) {
    cMap, maxCoord, _:= parse("input2.txt")
    for i := 0; i < b.N; i++ {
        part1(cMap, maxCoord)
    }
}
