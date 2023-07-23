package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func max() (int, error) {
	// Input will be a separate text file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	// The input data separtes each Elf's list by a new line then blank line
	var max int
	snackPool := strings.Split(string(data), "\n\n")
	for _, v := range snackPool {
		personalSnacks := strings.Split(v, "\n")
		var total int
		for _, v := range personalSnacks {
			num, err := strconv.Atoi(v)
			if err != nil {
				// The input data is bad and I can't be bothered
			}
			total += num
		}
		if total > max {
			max = total
		}
	}
	return max, err
}

func main() {
	max, err := max()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(max)
}
