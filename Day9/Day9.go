package main

import (
	"bufio"
	"fmt"
	"os"
)

func parse(fileName string) ([]int, error) {
	var parsedInput []int
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		intLine := make([]int, len(line))
		for i, char := range line {
			intLine[i] = int(char - '0')
		}
		filePointer, freePointer := 0, 1
		fileId := 0
		lenLine := len(intLine)
		for filePointer < lenLine || freePointer < lenLine {
			for i := 0; filePointer < lenLine && i < intLine[filePointer]; i++ {
				parsedInput = append(parsedInput, fileId)
			}
			for i := 0; freePointer < lenLine && i < intLine[freePointer]; i++ {
				parsedInput = append(parsedInput, -1)
			}
			filePointer += 2
			freePointer += 2
			fileId += 1
		}
	}
	return parsedInput, nil
}

func checkSum(input []int) int {
	checkSum := 0
	for i := 0; i < len(input); i++ {
		if input[i] != -1 {
			checkSum += i * input[i]
		}
	}
	return checkSum
}

func part2(input []int) int {
	lenInput := len(input)
	left, rightLeft, rightRight := 0, lenInput - 1, lenInput - 1
	for left < rightLeft {
		for left < rightLeft && input[left] != -1 {
			left += 1
		}
		for rightRight > left && input[rightRight] == -1 {
			rightRight--
		}
		rightLeft = rightRight
		for rightLeft-1 > left && input[rightLeft-1] == input[rightRight] {
			rightLeft--
		}
		windowLength := rightRight - rightLeft
		swapped := false
		for left < rightLeft && !swapped {
			spaceAvailable := true
			for i := left; i <= left + windowLength; i++ {
				if input[i] != -1 {
					spaceAvailable = false
					break
				}
			}
			if spaceAvailable {
				for i := 0 ; i <= windowLength ; i++ {
					input[left + i], input[rightRight - i] = input[rightRight - i], input[left + i]
				}
				swapped = true
			} else {
				left++
				for left < rightLeft && input[left] != -1 {
					left++
				}
			}
		}
		rightRight, rightLeft = rightLeft - 1, rightLeft - 1
		left = 0
	}
	return checkSum(input)
}

func part1(input []int) int {
	left := 0
	right := len(input) - 1
	for left < right {
		for left < len(input) && input[left] != -1 {
			left += 1
		}
		for right >= 0 && input[right] == -1 {
			right -= 1
		}
		if left < right {
			input[left], input[right] = input[right], input[left]
			left++
			right--
		}
	}
	return checkSum(input)
}

func main() {
	parsedInp, err := parse("input2.txt")
	if err != nil {
		return
	}
	part1 := part1(append([]int(nil), parsedInp...))
	fmt.Println("Part 1 : ", part1)
	part2 := part2(append([]int(nil), parsedInp...))
	fmt.Println("Part 2 : ", part2)
}
