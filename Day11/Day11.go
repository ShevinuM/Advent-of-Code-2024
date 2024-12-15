package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

var	stones = make(map[int]int)

func intPow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func parse(filename string) map[int]int {
	stones = make(map[int]int)
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			stones[num]++
		}
	}
	return stones
}

func getNumDigits(x int) int {
	num := 0
	for temp := x; temp > 0; temp /= 10 {
		num += 1
	}
	return num
}

func blink() {
	newStones := make(map[int]int)
	for stone, c := range stones {
		numDigits := getNumDigits(stone)
		if stone == 0 {
			newStones[1] += c
		} else if numDigits % 2 == 0 {
			div := intPow(10, numDigits / 2)
			s1 := stone / div
			s2 := stone % div
			newStones[s1] += c
			newStones[s2] += c
		} else {
			newStones[stone*2024] += c
		}
	}
	stones = newStones
}

func sol(iterations int) int {
	for i := 0; i < iterations; i++ {
		blink()
	}
	sum := 0
	for _, c := range stones {
		sum += c
	}
	return sum
}

func main() {
	parse("input2.txt")
	part1 := sol(25)
	fmt.Println("Part 1 : ", part1)
	part2 := sol(50)
	fmt.Println("Part 2 : ", part2)
}
