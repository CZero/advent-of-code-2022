// https://adventofcode.com/2022/day/1
// Day 1: Calorie Counting
package main

import (
	"aoc/libaoc"
	"fmt"
)

func main() {
	// lines, err := libaoc.ReadLines("example.txt")
	lines, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("Ruh Roh: No input")
	}
	fmt.Printf("Most calories found: %d\n", mostCalories(lines))
}

// mostCalories counts all elves' calories and reports the highest number
func mostCalories(lines []string) int {
	var (
		currentElf      int
		highestCalories int
	)
	for _, line := range lines {
		switch line {
		case "":
			if currentElf > highestCalories { // This one carries more
				highestCalories = currentElf
			}
			currentElf = 0 // Reset the count for the next elf
		default:
			calories := libaoc.SilentAtoi(line)
			currentElf += calories
		}
	}
	return highestCalories
}
