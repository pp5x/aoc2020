package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func countAreaTrees(area []string) int {
	count := 0
	j := 0
	for _, line := range area[1:] {
		j = (j + 3) % len(line)
		if line[j] == '#' {
			count++
		}
	}

	return count
}

func main() {
	buffer, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(buffer), "\n")
	// Last '\n' does not have content
	area := lines[:len(lines)-1]

	count := countAreaTrees(area)
	fmt.Println(count)
}
