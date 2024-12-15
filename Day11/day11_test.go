package main

import (
    "testing"
)

func BenchmarkPart1(b *testing.B) {
    parse("input2.txt")
    for i := 0; i < b.N; i++ {
        sol(25)
    }
}

func BenchmarkPart2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        sol(50)
    }
}