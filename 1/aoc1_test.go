package main

import "testing"

type Fixture struct {
	input    []int
	expected int
}

func TestCompute2020(t *testing.T) {
	fixtures := []Fixture{
		{[]int{1010}, -1},
		{[]int{2019, 1}, 2019},
		{[]int{1721, 299}, 514579},
		{[]int{1721, 979, 366, 299, 675, 1456}, 514579},
	}

	for _, fixture := range fixtures {
		res := compute2020(fixture.input)
		if res != fixture.expected {
			t.Error(res, fixture.expected)
		}
	}
}

func TestCompute2020Three(t *testing.T) {
	fixtures := []Fixture{
		{[]int{1, 2, 3}, -1},
		{[]int{979, 366, 675}, 241861950},
		{[]int{1721, 979, 366, 299, 675, 1456}, 241861950},
	}

	for _, fixture := range fixtures {
		res := compute2020Three(fixture.input)
		if res != fixture.expected {
			t.Error(res, fixture.expected)
		}
	}
}
