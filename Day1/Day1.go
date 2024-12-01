package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var input string
	input = "input2.txt"
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error in opening the file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int
	var leftList []int
	var rightList []int
	for scanner.Scan() {
		line := scanner.Text()
		lists := strings.Fields(line)
		if len(lists) == 2 {
			leftValue, err1 := strconv.Atoi(lists[0])
			rightValue, err2 := strconv.Atoi(lists[1])
			fmt.Println("Left Value : ", leftValue)
			fmt.Println("Right Value : ", rightValue)
			if err1 == nil && err2 == nil {
				leftList = append(leftList, leftValue)
				rightList = append(rightList, rightValue)
			}
		}

	}

	sort.Slice(leftList, func(i, j int) bool {
		return leftList[i] > leftList[j]
	})

	sort.Slice(rightList, func(i, j int) bool {
		return rightList[i] > rightList[j]
	})

	for len(leftList) > 0 && len(rightList) > 0 {
		lastLeft := leftList[(len(leftList) - 1)]
		leftList = leftList[:len(leftList)-1]
		lastRight := rightList[(len(rightList) - 1)]
		rightList = rightList[:len(rightList)-1]
		difference := math.Abs(float64(lastLeft - lastRight))
		sum += int(difference)
	}
	fmt.Println("Sum : ", sum)

}
