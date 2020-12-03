package main

import (
	"strings"
	"testing"
)

type Fixture struct {
	rightIncr int
	downIncr  int
	expected  int
}

func TestIsPasswordValidPolicy1(t *testing.T) {
	fixtures := []Fixture{
		{1, 1, 2},
		{3, 1, 7},
		{5, 1, 3},
		{7, 1, 4},
		{1, 2, 2},
	}

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
	for _, fixture := range fixtures {
		count := countAreaTrees(area, fixture.rightIncr, fixture.downIncr)

		if count != fixture.expected {
			t.Error(count, fixture.expected)
		}
	}
}
