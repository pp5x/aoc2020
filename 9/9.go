package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const MAX_UINT = ^uint(0)
const MAX_INT = int(MAX_UINT >> 1)
const MIN_INT = -MAX_INT - 1

func buildCandidatesTable(preamble []int) (candidates []int) {
	candidates = make([]int, 0, len(preamble)*len(preamble))

	for _, i := range preamble {
		for _, j := range preamble {
			candidates = append(candidates, i+j)
		}
	}

	return
}

// Unique algorithm on sorted array
func uniqueInts(ints []int) []int {
	if len(ints) <= 1 {
		return ints
	}

	prev := ints[0]
	end := 1

	for _, val := range ints {
		if val != prev {
			ints[end] = val
			end++
		}
		prev = val
	}

	return ints[:end]
}

func convertToInts(strs []string) ([]int, error) {
	ints := make([]int, len(strs))

	for i, str := range strs {
		if val, err := strconv.Atoi(str); err == nil {
			ints[i] = val
		} else {
			return nil, err
		}

	}

	return ints, nil
}

func isSummable(preamble []int, num int) bool {
	for _, val := range preamble {
		var candidate int
		if num > val {
			candidate = num - val
		} else {
			candidate = val - num
		}

		for _, val2 := range preamble {
			if candidate == val2 {
				return true
			}
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func minMaxSummable(preamble []int, num int) (int, int) {
	for i, _ := range preamble {
		sum := 0
		minV := MAX_INT
		maxV := -1
		for j := i; sum < num && j < len(preamble); j++ {
			sum += preamble[j]
			minV = min(minV, preamble[j])
			maxV = max(maxV, preamble[j])

			if sum == num {
				return minV, maxV
			}
		}
	}

	return -1, -1
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	buffer, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(MIN_INT)

	str := strings.TrimSpace(string(buffer))
	input, _ := convertToInts(strings.Split(str, "\n"))
	const PreambleSize int = 25
	for i := PreambleSize; i < len(input); i++ {
		if !isSummable(input[i-PreambleSize:i], input[i]) {
			min, max := minMaxSummable(input, input[i])
			fmt.Println(input[i])
			fmt.Println(min, max)
			fmt.Println(min + max)
		}
	}
}
