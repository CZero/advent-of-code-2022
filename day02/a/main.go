// https://adventofcode.com/2022/day/2
// Day 2: Rock Paper Scissors

// A rock  1
// B Paper 2
// C Scissors 3
// 0 lost
// 3 draw
// 6 win

package main

import (
	"aoc/libaoc"
	"fmt"
	"strings"
)

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	fmt.Println(playGuide(input))
}

func playGuide(input []string) int {
	var score int
	for _, line := range input {
		hands := strings.Fields(line)
		score += calcRound(hands)
	}
	return score
}

func calcRound(hands []string) (score int) {
	mappedHands := map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
		"Y": "paper",
		"X": "rock",
		"Z": "scissors",
	}
	switch mappedHands[hands[0]] {
	case "rock":
		switch mappedHands[hands[1]] {
		case "rock":
			return 4
		case "paper":
			return 8
		case "scissors":
			return 3
		}
	case "paper":
		switch mappedHands[hands[1]] {
		case "rock":
			return 1
		case "paper":
			return 5
		case "scissors":
			return 9
		}
	case "scissors":
		switch mappedHands[hands[1]] {
		case "rock":
			return 7
		case "paper":
			return 2
		case "scissors":
			return 6
		}
	default:
		panic("Shouldn't reach this!")
	}
	return 0
}
