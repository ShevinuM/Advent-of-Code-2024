package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func read() ([][]int, error) {

	file, err := os.Open("input2.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	var lines [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		var intFields []int
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			intFields = append(intFields, num)
		}
		lines = append(lines, intFields)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	return lines, nil
}

func part1(lines [][]int) (int, error) {
	sum := 0
	for i := 0; i < len(lines); i++ {
		isDescending := lines[i][0] > lines[i][1]
		isValid := true
		for j := 0; j < len(lines[i])-1; j++ {
			difference := math.Abs(float64(lines[i][j] - lines[i][j+1]))
			if difference < 1 || difference > 3 {
				isValid = false
				break
			}
			if isDescending && lines[i][j] > lines[i][j+1] {
				continue
			} else if !isDescending && lines[i][j] < lines[i][j+1] {
				continue
			} else {
				isValid = false
				break
			}
		}
		if isValid {
			sum += 1
		}
	}
	return sum, nil
}

func main() {
	lines, err := read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sum, err := part1(lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Part 1 -> %d\n", sum)
}
