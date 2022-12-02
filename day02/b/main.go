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
	"fmt"
	"oac/libaoc"
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
		"Y": "draw",
		"X": "loose",
		"Z": "win",
	}
	switch mappedHands[hands[0]] {
	case "rock":
		switch mappedHands[hands[1]] {
		case "draw": // rock
			return 3 + 1
		case "loose": // scissors
			return 0 + 3
		case "win": // paper
			return 6 + 2
		}
	case "paper":
		switch mappedHands[hands[1]] {
		case "draw": // paper
			return 3 + 2
		case "loose": // rock
			return 0 + 1
		case "win": // scissors
			return 6 + 3
		}
	case "scissors":
		switch mappedHands[hands[1]] {
		case "draw": // scissors
			return 3 + 3
		case "loose": // paper
			return 0 + 2
		case "win": // stone
			return 6 + 1
		}
	default:
		panic("Shouldn't reach this!")
	}
	return 0
}
