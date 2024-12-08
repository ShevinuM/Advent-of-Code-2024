package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"gonum.org/v1/gonum/stat/combin"
)

type Coordinate struct {
	y int
	x int
}

type Set map[Coordinate]struct{}

func isAlphaNumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

func parse(filename string) (map[rune][]Coordinate, Coordinate, error) {
	coordinateMap := make(map[rune][]Coordinate)
	file, err := os.Open(filename)
	if err != nil {
		return nil, Coordinate{y: -1, x: -1}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var maxX int
	var maxY int
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := []rune(line)
		maxX = len(parts) - 1
		for x := 0; x < len(parts); x++ {
			c := parts[x]
			if isAlphaNumeric(c) {
				if _, exists := coordinateMap[c]; exists {
					coordinateMap[c] = append(coordinateMap[c], Coordinate{y: y, x: x})
				} else {
					coordinateMap[c] = []Coordinate{Coordinate{y: y, x: x}}
				}
			}
		}
		y += 1
	}
	maxY = y - 1
	return coordinateMap, Coordinate{y: maxY, x: maxX}, nil
}

func isWithinBounds(coord Coordinate, maxCoord Coordinate, minCoord Coordinate) bool{
	return coord.y >= minCoord.y && coord.y <= maxCoord.y && coord.x >= minCoord.x && coord.x <= maxCoord.x
}

func sol(cMap map[rune][]Coordinate, maxCoord Coordinate, part int) int {
	minCoord := Coordinate{y: 0, x: 0}
	antinodes := make(Set) 
	for _, coords := range cMap {
		combs := combin.Combinations(len(coords), 2)
		for _, comb := range combs {
			var combi []Coordinate
			for _, index := range comb {
				combi = append(combi, coords[index])
			}
			yDiff := combi[1].y - combi[0].y
			xDiff := combi[1].x - combi[0].x
			antiNode1 := Coordinate{y: combi[1].y + yDiff, x: combi[1].x + xDiff}
			if isWithinBounds(antiNode1, maxCoord, minCoord) {
					antinodes[antiNode1] = struct{}{}
			}
			antiNode2 := Coordinate{y: combi[0].y - yDiff, x: combi[0].x - xDiff}
			if isWithinBounds(antiNode2, maxCoord, minCoord) {
					antinodes[antiNode2] = struct{}{}
			}
			if part == 2 {
				antiNode3 := Coordinate{y: combi[1].y - yDiff, x: combi[1].x - xDiff}
				for isWithinBounds(antiNode3, maxCoord, minCoord) {
					antinodes[antiNode3] = struct{}{}
					antiNode3 = Coordinate{y: antiNode3.y - yDiff, x: antiNode3.x - xDiff}
				}
				antiNode4 := Coordinate{y: combi[0].y + yDiff, x: combi[0].x + xDiff}
				for isWithinBounds(antiNode4, maxCoord, minCoord) {
					antinodes[antiNode4] = struct{}{}
					antiNode4 = Coordinate{y: antiNode4.y + yDiff, x: antiNode4.x + xDiff}
				}
			}
		}
	}
	return len(antinodes)
}

func main() {
	cMap, maxCoord, err := parse("input2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	part1 := sol(cMap, maxCoord, 1)
	fmt.Println("Part 1 -> ", part1)
	part2 := sol(cMap, maxCoord, 2)
	fmt.Println("Part 2 -> ", part2)
}
