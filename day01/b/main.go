// https://adventofcode.com/2022/day/1
// Day 1: Calorie Counting
package main

import (
	"fmt"
	"oac/libaoc"
	"sort"
)

func main() {
	// lines, err := libaoc.ReadLines("example.txt")
	lines, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("Ruh Roh: No input")
	}
	var topScores []int
	topScores = mostCalories(lines, topScores)
	fmt.Printf("Most calories found: %d\n%v\n", sumScores(topScores), topScores)
}

// mostCalories counts all elves' calories and reports the highest three
func mostCalories(lines []string, topScores []int) []int {
	var currentElf int
	for i, line := range lines {
		if line == "" { // Empty line, so the last elf is done
			topScores = addIfBigger(currentElf, topScores)
			currentElf = 0 // Reset the count for the next elf
		} else {
			calories := libaoc.SilentAtoi(line)
			currentElf += calories
			if i == len(lines)-1 { // If no more lines are coming, also tally the last elf
				topScores = addIfBigger(currentElf, topScores)
			}
		}
	}
	return topScores
}

// sortTopscores sorts the []int topScores
func sortTopscores(topScores []int) []int {
	sort.Ints(topScores)
	return topScores
}

// addIfBigger adds the score if it's bigger than the numbers present.
func addIfBigger(score int, topScores []int) []int {
	if len(topScores) < 3 { // Less than 3 scores, we can add it anyway.
		topScores = append(topScores, score)
		topScores = sortTopscores(topScores) // Sort scores to keep them organized
		return topScores
	}
	for i := 2; i >= 0; i-- { // We have three scores. We'll compare from right to left.
		if score > topScores[i] {
			for j := 0; j < i; j++ { // We found one that's bigger. We shift the rest left.
				topScores[j] = topScores[j+1]
			}
			topScores[i] = score
			return topScores
		}
	}
	return topScores
}

// sumScores sums the scores in topScores []int
func sumScores(topScores []int) (total int) {
	for _, n := range topScores {
		total += n
	}
	return total
}
