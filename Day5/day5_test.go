package main

import (
    "testing"
)

func BenchmarkPart1(b *testing.B) {
    po, pl, _ := read("input2.txt")
    for i := 0; i < b.N; i++ {
        part1(po, pl)
    }
}

func BenchmarkPart2(b *testing.B) {
    po, pl, por := read("input2.txt")
    for i := 0; i < b.N; i++ {
        part2(po, pl, por)
    }
}