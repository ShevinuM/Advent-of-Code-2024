package main

import (
    "testing"
)

func BenchmarkPart1(b *testing.B) {
    parsedInp, _ := parse("input2.txt")
    for i := 0; i < b.N; i++ {
        part1(parsedInp)
    }
}

func BenchmarkPart2(b *testing.B) {
    parsedInp, _ := parse("input2.txt")
    for i := 0; i < b.N; i++ {
        part2(parsedInp)
    }
}