package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func read() ([]int, error) {
	// Input will be a separate text file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		return nil, err
	}
	// The input data separtes each Elf's list by a new line then blank line
	var snacks []int
	snacksPool := strings.Split(string(data), "\n\n")
	for _, v := range snacksPool {
		var personalTotal int
		personalPool := strings.Split(v, "\n")
		for _, v := range personalPool {
			itemCalories, err := strconv.Atoi(v)
			if err != nil {
				// The input data is bad (trailing empty line) and I can't be bothered
			}
			personalTotal += itemCalories
		}
		snacks = append(snacks, personalTotal)
	}
	return snacks, err
}

func main() {
	snacks, err := read()
	if err != nil {
		log.Fatal(err)
	}
	sort.Ints(snacks)
	top3 := snacks[len(snacks)-3:]
	var total int
	for _, v := range top3 {
		total += v
	}
	fmt.Println(total)
	// fmt.Println(snacks[len(snacks)-3:])
}
