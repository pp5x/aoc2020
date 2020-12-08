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

func unloop(program []string, executed []bool, pc int, patched bool) (bool, int) {
	acc := 0
	for pc < len(program) {
		if executed[pc] {
			fmt.Println(acc)
			return false, acc
		}

		tokens := strings.Split(program[pc], " ")

		switch tokens[0] {
		case "nop":
			if !patched {
				programPatched := make([]string, len(program))
				copy(programPatched, program)
				executedCopy := make([]bool, len(executed))
				programPatched[pc] = strings.Replace(programPatched[pc], "nop", "jmp", 1)
				endReached, accRec := unloop(programPatched, executedCopy, pc, true)
				if endReached {
					return true, acc + accRec
				}
			}
			executed[pc] = true
			pc++
		case "acc":
			val, _ := strconv.Atoi(tokens[1])
			acc += val
			executed[pc] = true
			pc++
		case "jmp":
			if !patched {
				programPatched := make([]string, len(program))
				copy(programPatched, program)
				executedCopy := make([]bool, len(executed))
				programPatched[pc] = strings.Replace(programPatched[pc], "jmp", "nop", 1)
				endReached, accRec := unloop(programPatched, executedCopy, pc, true)
				if endReached {
					return true, acc + accRec
				}
			}
			val, _ := strconv.Atoi(tokens[1])
			executed[pc] = true
			pc += val
		}
	}

	return true, acc
}

func main() {
	buffer, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	str := strings.TrimSpace(string(buffer))
	program := strings.Split(str, "\n")
	executed := make([]bool, len(program))

	fmt.Println(unloop(program, executed, 0, false))
}
