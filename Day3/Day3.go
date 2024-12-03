package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func read() string{
	file, err := os.Open("input2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	return string(content)
}

func part1(input string) int {
	sum := 0
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		if len(match) == 3 {
			X := match[1]
			Y := match[2]
			x, err := strconv.Atoi(X)
			if err != nil {
				fmt.Println(err)
			}
			y, err := strconv.Atoi(Y)
			if err != nil {
				fmt.Println(err)
			}
			sum += x * y
		}
	}
	return sum
}

func main() {
	content := read()
	part1 := part1(content)
	fmt.Println("Part 1: ", part1)
}
