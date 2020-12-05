package main

import (
	"testing"
)

type Fixture struct {
	boardingPass string
	expected     int
}

func TestGetSeatRow(t *testing.T) {
	rowFixtures := []Fixture{
		{"F", 0},
		{"FF", 0},
		{"FFFFFFF", 0},
		{"B", 64},
		{"BB", 96},
		{"BBBBBBB", 127},
		{"FBFBBFF", 44},
		{"BFFFBBF", 70},
		{"FFFBBBF", 14},
		{"BBFFBBF", 102},
	}

	for _, fixture := range rowFixtures {
		seatID := getSeat(fixture.boardingPass, 'F', 128)
		if seatID != fixture.expected {
			t.Error(seatID, fixture.expected)
		}
	}

	seatID := getSeat("RLR", 'L', 8)
	if seatID != 5 {
		t.Error(seatID, 5)
	}

	seatID = getSeat("RRR", 'L', 8)
	if seatID != 7 {
		t.Error(seatID, 7)

	}

	seatID = getSeat("RLL", 'L', 8)
	if seatID != 4 {
		t.Error(seatID, 4)
	}
}

func TestGetSeatID(t *testing.T) {
	fixtures := []Fixture{
		{"FBFBBFFRLR", 357},
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}

	for _, fixture := range fixtures {
		seatID := getSeatID(fixture.boardingPass)
		if seatID != fixture.expected {
			t.Error(seatID, fixture.expected)
		}
	}
}
