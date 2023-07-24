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
	loss := map[int]int{2: 1, 3: 2, 1: 3}
	var totalScore int
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Drop the last item as the last line ends in newline too
	r := strings.Split(string(dat), "\n")
	rounds := r[:len(r)-1]
	for _, v := range rounds {
		hand := strings.Split(v, " ")
		opponent := m[hand[0]]
		strat := m[hand[1]]
		if strat == 2 {
			totalScore += 3 + opponent
		} else if strat == 3 {
			totalScore += 6 + champion[opponent]
		} else {
			totalScore += 0 + loss[opponent]
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
