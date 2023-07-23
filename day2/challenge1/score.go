package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func get() (int, error) {
	m := map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}
	// key is opponent move. value is the move to beat them. compare to your move to see if you win
	champion := map[int]int{1: 2, 2: 3, 3: 1}
	var totalScore int
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Drop the last item as the last line ends in newline too
	r := strings.Split(string(dat), "\n")
	rounds := r[:len(r)-1]
	for _, v := range rounds {
		hands := strings.Split(v, " ")
		opponent := m[hands[0]]
		you := m[hands[1]]
		if you == opponent {
			totalScore += 3 + you
		} else if you == champion[opponent] {
			totalScore += 6 + you
		} else {
			totalScore += 0 + you
		}
	}
	return totalScore, nil
}

func main() {
	total, err := get()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}
