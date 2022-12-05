// https://adventofcode.com/2022/day/5
// Day 5: Supply Stacks
package main

import (
	"fmt"
	"oac/libaoc"
	"strings"
)

type (
	Stack []string
	Move  struct {
		num  int
		from int
		to   int
	}
)

var stacks []Stack

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	layout, instructions := separateInput(input) // Separate the layout and the moves

	// Setup the stacks and the moves! (parse them into data)
	fillStacks(layout)
	moves := parseInstructions(instructions)

	moveCrates(moves)
	fmt.Printf("Resulting stacks: \n%v\n\n", stacks)
	fmt.Printf("First crates in the stacks:\n%s\n", getFirstCrates())
}

func moveCrates(moves []Move) {
	for _, move := range moves {
		crates := stacks[move.from][:move.num]
		stacks[move.from] = stacks[move.from][move.num:]
		var newstack Stack
		for _, crate := range crates {
			newstack = append(newstack, crate)
		}
		for _, crate := range stacks[move.to] {
			newstack = append(newstack, crate)
		}
		stacks[move.to] = newstack
	}
}

func getFirstCrates() (firstcrates string) {
	for _, stack := range stacks {
		firstcrates += stack[0]
	}
	return firstcrates
}

func fillStacks(layout []string) {
	stacks = make([]Stack, (len(layout[2])+1)/4)
	for i, _ := range stacks {
		stacks[i] = make(Stack, 0)
	}
	for i, line := range layout {
		if i == len(layout)-1 {
			continue
		}
		for j, crate := range line {
			if (j-1)%4 == 0 {
				if string(crate) != " " {
					stacks[((j - 1) / 4)] = append(stacks[((j-1)/4)], string(crate))
				}
			}
		}
	}
	return
}

func parseInstructions(instructions []string) (moves []Move) {
	for _, instruction := range instructions {
		parts := strings.Fields(instruction)
		moves = append(moves, Move{libaoc.SilentAtoi(parts[1]), libaoc.SilentAtoi(parts[3]) - 1, libaoc.SilentAtoi(parts[5]) - 1})
	}
	return moves
}

func separateInput(input []string) (layout, instructions []string) {
	var separatorSeen bool
	for _, line := range input {
		switch line {
		case "":
			separatorSeen = true
		default:
			if !separatorSeen {
				layout = append(layout, line)
			} else {
				instructions = append(instructions, line)
			}
		}
	}
	return layout, instructions
}
