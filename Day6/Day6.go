package main

import (
	"bufio"
	"fmt"
	"github.com/deckarep/golang-set/v2"
	"os"
)

type Coordinate struct {
	Y int
	X int
}

func parse(filename string) [][]rune {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var matrix [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := []rune(line)
		matrix = append(matrix, parts)
	}

	return matrix
}

func findStart(matrix [][]rune) (int, int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			character := matrix[i][j]
			if character == '^' || character == 'v' || character == '>' || character == '<' {
				return i, j
			}
		}
	}
	return -1, -1
}

func getMoveCoordinates(direction rune) []int {
	switch direction {
	case '^':
		return []int{-1, 0}
	case 'v':
		return []int{1, 0}
	case '>':
		return []int{0, 1}
	case '<':
		return []int{0, -1}
	default:
		return []int{0, 0}
	}
}

func getNextDirection(curr rune) rune {
	switch curr {
	case '^':
		return '>'
	case 'v':
		return '<'
	case '>':
		return 'v'
	case '<':
		return '^'
	default:
		return curr
	}
}

func part1(matrix [][]rune) int {
	y, x := findStart(matrix)
	if y == -1 || x == -1 {
		return -1
	}
	finished := false
	visited := mapset.NewSet[Coordinate]()
	currDirec := matrix[y][x]
	moveCoords := getMoveCoordinates(matrix[y][x])
	for !finished {
		if y >= len(matrix) || y < 0 || x >= len(matrix[0]) || x < 0 {
			finished = true
			return visited.Cardinality()
		}
		if matrix[y][x] == '#' {
			y = y - moveCoords[0]
			x = x - moveCoords[1]
			nextDir := getNextDirection(currDirec)
			currDirec = nextDir
			moveCoords = getMoveCoordinates(currDirec)
		}
		coord := Coordinate{Y: y, X: x}
		matrix[y][x] = 'X'
		visited.Add(coord)
		y = y + moveCoords[0]
		x = x + moveCoords[1]
	}
	return visited.Cardinality()
}

func main() {
	parsedInput := parse("input2.txt")
	part1 := part1(parsedInput)
	fmt.Println(part1)
}
