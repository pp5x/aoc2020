package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func compute2020Rec(input []int, depth int, totalSum int, totalMul int) int {
	for i, val := range input {
		if depth <= 0 {
			if totalSum+val == 2020 {
				return totalMul * val
			}
		} else {
			res := compute2020Rec(input[i+1:], depth-1, totalSum+val, totalMul*val)
			if res > 0 {
				return res
			}
		}
	}

	return -1
}

func compute2020(input []int) int {
	return compute2020Rec(input, 1, 0, 1)
}

func compute2020Three(input []int) int {
	return compute2020Rec(input, 2, 0, 1)
}

func main() {
	buffer, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(buffer), "\n")
	// Last '\n' does not have content
	strValues := lines[:len(lines)-1]

	var numbers []int
	for _, value := range strValues {
		num, err := strconv.Atoi(value)

		numbers = append(numbers, num)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(compute2020(numbers))
	fmt.Println(compute2020Three(numbers))
}
