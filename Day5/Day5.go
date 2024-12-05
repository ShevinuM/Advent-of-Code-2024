package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read(filename string) (map[int][]int, [][]int, map[int][]int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Can't read file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	section1 := true
	pageOrder := make(map[int][]int)
	pageOrderReverse := make(map[int][]int)
	var pagesList [][]int
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			section1 = false
			continue
		}

		if section1 {
			parts := strings.Split(line, "|")
			if len(parts) == 2 {
				page1, err1 := strconv.Atoi(parts[0])
				page2, err2 := strconv.Atoi(parts[1])
				if err1 != nil || err2 != nil {
					fmt.Println("error in converting pages to int")
				}
				pageOrder[page1] = append(pageOrder[page1], page2)
				pageOrderReverse[page2] = append(pageOrderReverse[page2], page1)
			} else {
				fmt.Println("Invalid Pages")
			}
		} else {
			parts := strings.Split(line, ",")
			var pages []int
			for _, part := range parts {
				pg, err := strconv.Atoi(part)
				if err != nil {
					fmt.Println("Error in input")
				}
				pages = append(pages, pg)
			}
			pagesList = append(pagesList, pages)
		}
	}
	return pageOrder, pagesList, pageOrderReverse
}

func check(pageOrder map[int][]int, update []int) (bool, int) {
	for i := 0; i < len(update); i++ {
		val := update[i]
		for j := i; j < len(update); j++ {
			val2 := update[j]
			if postVals, exists := pageOrder[val2]; exists {
				for _, postVal := range postVals {
					if postVal == val {
						return false, j
					}
				}
			}
		}
	}
	return true, -1
}

func sort(pageOrder map[int][]int, update []int, pageOrderReverse map[int][]int) []int {
	for i := 0; i < len(update); i++ {
		val := update[i]
		for j := i + 1; j < len(update); j++ {
			val2 := update[j]
			if postVals, exists := pageOrder[val2]; exists {
			outerLoop:
				for _, postVal := range postVals {
					if postVal == val {
						for i2 := 0; i2 < j; i2++ {
							valPro := update[i2]
							if preVals, exists := pageOrderReverse[valPro]; exists {
								for _, preVal := range preVals {
									if preVal == val2 {
										removedValue := update[j]
										update = append(update[:j], update[j+1:]...)
										update = append(update[:i2], append([]int{removedValue}, update[i2:]...)...)
										break outerLoop
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return update
}
func part1(pageOrder map[int][]int, pages [][]int) int {
	sum := 0
	for _, update := range pages {
		valid, index := check(pageOrder, update)
		middleIndex := len(update) / 2
		middleElement := update[middleIndex]
		if valid || index == -1 {
			sum += middleElement
		}
	}
	return sum
}

func part2(pageOrder map[int][]int, pages [][]int, pageOrderReverse map[int][]int) int {
	sum := 0
	for _, update := range pages {
		valid, _ := check(pageOrder, update)
		middleIndex := len(update) / 2
		if !valid {
			newUpdate := sort(pageOrder, update, pageOrderReverse)
			fmt.Println(newUpdate)
			middleElement := newUpdate[middleIndex]
			sum += middleElement
		}
	}
	return sum
}

func main() {
	pageOrder, pagesList, pageOrderReverse := read("input2.txt")
	part1 := part1(pageOrder, pagesList)
	fmt.Println(part1)
	part2 := part2(pageOrder, pagesList, pageOrderReverse)
	fmt.Println(part2)
}
