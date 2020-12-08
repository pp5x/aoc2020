package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	buffer, err := ioutil.ReadFile("input_aoc.txt")
	if err != nil {
		log.Fatal(err)
	}

	str := strings.TrimSpace(string(buffer))

	acc := 0
	program := strings.Split(str, "\n")
	executed := make([]bool, len(program))
	for pc := 0; pc < len(program); {
		if executed[pc] {
			fmt.Println(acc)
			break
		}

		tokens := strings.Split(program[pc], " ")

		executed[pc] = true
		switch tokens[0] {
		case "nop":
			pc++
		case "acc":
			val, _ := strconv.Atoi(tokens[1])
			acc += val
			pc++
		case "jmp":
			val, _ := strconv.Atoi(tokens[1])
			pc += val
		}
	}
}
