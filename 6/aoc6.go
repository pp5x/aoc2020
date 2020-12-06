package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	buffer, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	sumCommon := 0
	chunks := strings.Split(string(buffer), "\n\n")
	for _, chunk := range chunks {
		unique := make(map[rune]int)
		chunk = strings.TrimSpace(chunk)
		nbPeople := len(strings.Split(chunk, "\n"))
		chunk = strings.ReplaceAll(chunk, "\n", "")

		for _, c := range chunk {
			if val, exists := unique[c]; !exists {
				unique[c] = 1
			} else {
				unique[c] = val + 1
			}
		}

		for k, v := range unique {
			if v == nbPeople {
				fmt.Printf("%c %d\n", k, v)
				sumCommon++
			}
		}
		sum += len(unique)

		fmt.Printf("--\n")
	}
	fmt.Println(sum)
	fmt.Println(sumCommon)
}
