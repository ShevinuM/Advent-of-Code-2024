package main

import (
	"bufio"
	"fmt"
	"os"
)

func read(input string) [][]rune {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		characters := []rune(line)
		lines = append(lines, characters)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return lines
}

func part1(lines [][]rune) int {
	sum := 0
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] == 'X' {
				if y+3 < len(lines) {
					if lines[y+1][x] == 'M' && lines[y+2][x] == 'A' && lines[y+3][x] == 'S' {
						sum += 1
					}
				}
				if y-3 >= 0 {
					if lines[y-1][x] == 'M' && lines[y-2][x] == 'A' && lines[y-3][x] == 'S' {
						sum += 1
					}
				}
				if x+3 < len(lines[y]) {
					if lines[y][x+1] == 'M' && lines[y][x+2] == 'A' && lines[y][x+3] == 'S' {
						sum += 1
					}
				}
				if x-3 >= 0 {
					if lines[y][x-1] == 'M' && lines[y][x-2] == 'A' && lines[y][x-3] == 'S' {
						sum += 1
					}
				}
				if y+3 < len(lines) && x+3 < len(lines[y]) {
					if lines[y+1][x+1] == 'M' && lines[y+2][x+2] == 'A' && lines[y+3][x+3] == 'S' {
						sum += 1
					}
				}
				if y+3 < len(lines) && x-3 >= 0 {
					if lines[y+1][x-1] == 'M' && lines[y+2][x-2] == 'A' && lines[y+3][x-3] == 'S' {
						sum += 1
					}
				}
				if y-3 >= 0 && x+3 < len(lines[y]) {
					if lines[y-1][x+1] == 'M' && lines[y-2][x+2] == 'A' && lines[y-3][x+3] == 'S' {
						sum += 1
					}
				}
				if y-3 >= 0 && x-3 >= 0 {
					if lines[y-1][x-1] == 'M' && lines[y-2][x-2] == 'A' && lines[y-3][x-3] == 'S' {
						sum += 1
					}
				}
			}
		}
	}
	return sum
}

func main() {
	lines := read("input2.txt")
	part1 := part1(lines)
	fmt.Println(part1)
}
