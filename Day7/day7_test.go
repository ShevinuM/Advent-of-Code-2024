package main

import (
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	parsedInput, _ := parse("input2.txt")
	for i := 0; i < b.N; i++ {
		part1(parsedInput)
	}
}

func BenchmarkPart2(b *testing.B) {
	parsedInput, _ := parse("input2.txt")
	for i := 0; i < b.N; i++ {
		part2(parsedInput)
	}
}