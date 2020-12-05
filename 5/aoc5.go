package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func getSeat(boardingPass string, divideCharacter byte, maxSeatRow int) int {
	if boardingPass == "" {
		return 0
	}

	if boardingPass[0] == divideCharacter {
		return getSeat(boardingPass[1:], divideCharacter, maxSeatRow>>1)
	}

	return getSeat(boardingPass[1:], divideCharacter, maxSeatRow>>1) + maxSeatRow>>1
}

func getSeatID(boardingPass string) int {
	return getSeat(boardingPass[:7], 'F', 128)*8 + getSeat(boardingPass[7:], 'L', 8)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	buffer, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(buffer), "\n")
	lines = lines[:len(lines)-1]

	maxSeatID := 0
	minSeatID := 1024
	// Find missing ID by XORing all seats values.
	missingID := 0
	for _, line := range lines {
		seatID := getSeatID(line)
		maxSeatID = max(maxSeatID, seatID)
		minSeatID = min(minSeatID, seatID)
		missingID ^= seatID
	}

	for i := 0; i < minSeatID; i++ {
		missingID ^= i
	}
	for i := maxSeatID + 1; i < 1024; i++ {
		missingID ^= i
	}

	fmt.Println(missingID)
}
