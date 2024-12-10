package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set map[Coordinate]struct{}

type Coordinate struct {
	y   int
	x   int
	val int
}

func parse(filename string) ([][]int, []Coordinate, error) {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var trailHeads []Coordinate
	var topoMap [][]int
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")
		var nums []int
		for i, part := range parts {
			num, _ := strconv.Atoi(part)
			nums = append(nums, num)
			if num == 0 {
				trailHeads = append(trailHeads, Coordinate{y: y, x: i})
			}
		}
		topoMap = append(topoMap, nums)
		y++
	}
	return topoMap, trailHeads, nil
}

func bfs(topoMap [][]int, trailHead Coordinate, part int) int {
	lenY := len(topoMap)
	lenX := len(topoMap[0])
	sum := 0
	var queue []Coordinate
	visited := make(Set)
	queue = append(queue, trailHead)
	directions := []Coordinate{{y: 0, x: 1, val: 0}, {y: 0, x: -1, val: 0}, {y: 1, x: 0, val: 0}, {y: -1, x: 0, val: 0}}
	for len(queue) > 0 {
		trailPoint := queue[0]
		queue = queue[1:]
		for _, direction := range directions {
			newY := trailPoint.y + direction.y
			newX := trailPoint.x + direction.x
			newVal := trailPoint.val + 1
			newCoordinate := Coordinate{y: newY, x: newX, val: newVal}
			if newY >= 0 && newY < lenY && newX >= 0 && newX < lenX && newVal == 9 && newVal == topoMap[newY][newX] {
				if part == 1 && _, ok := visited[newCoordinate]; !ok {
					if _, ok := visited[newCoordinate]; !ok {
						sum++
					}
				} else {
					sum++
				}
				visited[newCoordinate] = struct{}{}
			} else if newY >= 0 && newY < lenY && newX >= 0 && newX < lenX && topoMap[newY][newX] == newVal {
				trailPoint := Coordinate{y: newY, x: newX, val: newVal}
				queue = append(queue, trailPoint)
			}
		}
	}
	return sum
}

func sol(topoMap [][]int, trailHeads []Coordinate, part int) int {
	sum := 0
	for _, trailHead := range trailHeads {
		val := bfs(topoMap, trailHead, part)
		sum += val
	}
	return sum
}

func main() {
	topoMap, trailHeads, _ := parse("input2.txt")
	part1 := sol(topoMap, trailHeads, 1)
	fmt.Println("Part 1 : ", part1)
	part2 := sol(topoMap, trailHeads, 2)
	fmt.Println("Part 2 : ", part2)
}
