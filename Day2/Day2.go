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
		line := lines[i]
		isValid, err := check1(line)
		if err != nil {
			os.Exit(1)
		}
		if isValid {
			sum += 1
		}
	}
	return sum, nil
}

func part2(lines [][]int) (int, error) {
	sum := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		isValid, err := check1(line)
		if err != nil {
			os.Exit(1)
		}
		if isValid {
			sum += 1
		} else {
			isValid, err := check2(line)
			if err != nil {
				os.Exit(1)
			}
			if isValid {
				sum += 1
			}
		}
	}
	return sum, nil
}

func check2(line []int) (bool, error) {
	for i := 0; i < len(line); i++ {
		line2 := removeElement(append([]int(nil), line...), i)
		isValid, err := check1(line2)
		if err != nil {
			os.Exit(1)
		}
		if isValid {
			return true, nil
		}
	}
	return false, nil
}

func check1(line []int) (bool, error) {
	isDescending := line[0] > line[1]
	isValid := true
	for j := 0; j < len(line)-1; j++ {
		difference := math.Abs(float64(line[j] - line[j+1]))
		if difference < 1 || difference > 3 {
			isValid = false
			break
		}
		if isDescending && line[j] > line[j+1] {
			continue
		} else if !isDescending && line[j] < line[j+1] {
			continue
		} else {
			isValid = false
			break
		}
	}
	return isValid, nil
}

func removeElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
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
	sum2, err := part2(lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Part 2 -> %d\n", sum2)
}
