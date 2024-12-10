package main

import (
    "testing"
)

func BenchmarkPart1(b *testing.B) {
    tm, th, _ := parse("input2.txt")
    for i := 0; i < b.N; i++ {
        sol(tm, th, 1)
    }
}

func BenchmarkPart2(b *testing.B) {
    tm, th, _:= parse("input2.txt")
    for i := 0; i < b.N; i++ {
        sol(tm, th, 2)
    }
}