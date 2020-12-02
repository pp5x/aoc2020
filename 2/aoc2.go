package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func isPasswordValid(min int, max int, c byte, password []byte) bool {
	count := 0
	for _, character := range password {
		if character == c {
			count++
		}
	}

	return count >= min && count <= max
}

func isPasswordValid2(pos int, pos2 int, c byte, password []byte) bool {
	passwordLen := len(password)
	if pos > passwordLen || pos2 > passwordLen {
		panic("out of bound positions")
	}

	return (password[pos] == c) != (password[pos2] == c)

}

func main() {
	buffer, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(buffer), "\n")
	// Last '\n' does not have content
	strValues := lines[:len(lines)-1]

	validPasswordCountPolicy1 := 0
	validPasswordCountPolicy2 := 0

	for _, value := range strValues {
		tokens := strings.Split(value, " ")

		strCountPolicies := strings.Split(tokens[0], "-")
		characterPolicy := tokens[1][0]

		password := tokens[2]

		policyVal1, _ := strconv.Atoi(strCountPolicies[0])
		policyVal2, _ := strconv.Atoi(strCountPolicies[1])

		if isPasswordValid(policyVal1, policyVal2, characterPolicy, []byte(password)) {
			validPasswordCountPolicy1++
		}
		if isPasswordValid2(policyVal1-1, policyVal2-1, characterPolicy, []byte(password)) {
			validPasswordCountPolicy2++
		}
	}

	fmt.Println(validPasswordCountPolicy1)
	fmt.Println(validPasswordCountPolicy2)
}
