package main

import (
	"strings"
	"testing"
)

type Fixture struct {
	policyVal1 int
	policyVal2 int
	c          byte
	password   []byte
	expected   bool
}

func TestIsPasswordValidPolicy1(t *testing.T) {
	const rawArea = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

	area := strings.Split(rawArea, "\n")
	count := countAreaTrees(area)

	if count != 7 {
		t.Error(count, 7)
	}
}
