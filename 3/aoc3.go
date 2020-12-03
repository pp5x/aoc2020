package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func countAreaTrees(area []string, rightIncr, downIncr int) (count int) {
	i := 0
	for j := downIncr; j < len(area); j += downIncr {
		i = (i + rightIncr) % len(area[j])
		if area[j][i] == '#' {
			count++
		}
	}

	return count
}

type Slope struct {
	rightIncr int
	downIncr  int
}

func main() {
	buffer, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(buffer), "\n")
	// Last '\n' does not have content
	area := lines[:len(lines)-1]

	slopes := []Slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	res := 1
	for _, slope := range slopes {
		res *= countAreaTrees(area, slope.rightIncr, slope.downIncr)
	}

	fmt.Println(res)
}
