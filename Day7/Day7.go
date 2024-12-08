package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func parse(filename string) (map[int]map[int][]int, error) {
	equations := make(map[int]map[int][]int)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	r := bufio.NewReader(file)
	eof := false
	count := 0
	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF{
			return nil, err
		}
		if err == io.EOF {
			eof = true
		}
		equation := strings.Split(line, ":")
		if len(equation) != 2 {
			return nil, errors.New("Length of equation array not 2")
		}
		numsStr := strings.Fields(equation[1])
		nums := make([]int, len(numsStr))
		for i, num := range numsStr {
			nums[i], err = strconv.Atoi(num)
			if err != nil {
				return nil, err
			}
		}
		testVal, err := strconv.Atoi(equation[0])
		if err != nil {
			return nil, err
		}
		if equations[count] == nil {
			equations[count] = make(map[int][]int)
		}
		equations[count][testVal] = nums
		count += 1
		if eof {
			break
		}
	}
	return equations, nil
}

func calc(vals []int, operator rune, currSum int, targetSum int) bool {
	if len(vals) == 0 {
		return currSum == targetSum
	}
	x, vals := vals[0], vals[1:]
	if operator == '+' {
		currSum += x
	} else {
		currSum *= x
	}
	return calc(append([]int(nil), vals...), '+', currSum, targetSum) || calc(append([]int(nil), vals...), '*', currSum, targetSum)
}

func calc2(vals []int, operator rune, currSum int, targetSum int) bool {
	if len(vals) == 0 {
		return currSum == targetSum
	}
	x, vals := vals[0], vals[1:]
	if operator == '+' {
		currSum += x
	} else if operator == '|' {
		strA := strconv.Itoa(currSum)
		strB := strconv.Itoa(x)
		str := strA + strB
		currSum, _ = strconv.Atoi(str)
	} else {
		currSum *= x
	}
	return calc2(append([]int(nil), vals...), '+', currSum, targetSum) || calc2(append([]int(nil), vals...), '*', currSum, targetSum) || calc2(append([]int(nil), vals...), '|', currSum, targetSum) 
}

func part2(input map[int]map[int][]int) int {
	sum := 0
	for _, val := range input {
		for testVal, numbers := range val {
			isValid := calc2(append([]int(nil), numbers[1:]...), '+', numbers[0], testVal) || calc2(append([]int(nil), numbers[1:]...), '*', numbers[0], testVal) || calc2(append([]int(nil), numbers[1:]...), '|', numbers[0], testVal)
			if isValid {
				sum += testVal
			}
		}
	}
	return sum
}

func part1(input map[int]map[int][]int) int {
	sum := 0
	for _, val := range input {
		for testVal, numbers := range val {
			isValid := calc(append([]int(nil), numbers[1:]...), '+', numbers[0], testVal) || calc(append([]int(nil), numbers[1:]...), '*', numbers[0], testVal)
			if isValid {
				sum += testVal
			}
		}
	}
	return sum
}

func main() {
	parsedInput, err := parse("input2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	part1 := part1(parsedInput)
	fmt.Println("Part 1 : ", part1)
	part2 := part2(parsedInput)
	fmt.Println("Part 2 : ", part2)
}
