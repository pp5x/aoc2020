package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type BagStorage struct {
	name     string
	quantity int
}

func depthFirst(graph map[string][]BagStorage, bag string) int {
	count := 0
	for _, childBag := range graph[bag] {
		if childBag.name == "shiny gold" {
			return 1
		}

		count += depthFirst(graph, childBag.name)
	}

	return count
}

func depthFirstShinyGold(graph map[string][]BagStorage, bag string) int {
	count := 0
	for _, childBag := range graph[bag] {
		count += childBag.quantity
		count += childBag.quantity * depthFirstShinyGold(graph, childBag.name)
	}

	return count
}

func main() {
	buffer, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	str := strings.TrimSpace(string(buffer))

	graph := make(map[string][]BagStorage)
	for _, line := range strings.Split(str, "\n") {
		words := strings.Split(line, " ")

		var storageList []BagStorage
		for i := 4; i < len(words); i += 4 {
			q := 0
			if words[i] != "no" {
				q, _ = strconv.Atoi(words[i])
			}
			storageList = append(storageList, BagStorage{
				name:     strings.Join(words[i+1:i+3], " "),
				quantity: q,
			})
		}

		graph[strings.Join(words[0:2], " ")] = storageList
	}

	count := 0
	for bagName := range graph {
		if depthFirst(graph, bagName) > 0 {
			count++
		}
	}

	fmt.Println(count)
	fmt.Println(depthFirstShinyGold(graph, "shiny gold"))
}
