// https://adventofcode.com/2022/day/4
// Day 4: Camp Cleanup
package main

import (
	"fmt"
	"oac/libaoc"
	"strings"
)

type section struct {
	start int
	stop  int
}

type elf struct {
	s section
}

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	found, pairs := countOverlappingPairs(input)
	fmt.Printf("Pairs found:\n%v\n%d\n", pairs, found)
}

func countOverlappingPairs(input []string) (found int, pairs []string) {
	for _, pair := range input {
		if findContainedSections(pair) {
			pairs = append(pairs, pair)
			found++
		}
	}
	return found, pairs
}

func findContainedSections(pair string) bool {
	var elves []elf
	sections := strings.Split(pair, ",")
	for i, section := range sections {
		elves = append(elves, elf{})
		step := strings.Split(section, "-")
		elves[i].s.start, elves[i].s.stop = libaoc.SilentAtoi(step[0]), libaoc.SilentAtoi(step[1])
	}
	if isContained(elves) {
		return true
	}
	return false
}

func isContained(elves []elf) bool {
	switch {
	case elves[0].s.start >= elves[1].s.start && elves[0].s.start <= elves[1].s.stop:
		return true
	case elves[0].s.stop >= elves[1].s.start && elves[0].s.stop <= elves[1].s.stop:
		return true
	case elves[1].s.start >= elves[0].s.start && elves[1].s.start <= elves[0].s.stop:
		return true
	case elves[1].s.stop >= elves[0].s.start && elves[1].s.stop <= elves[0].s.stop:
		return true
	}
	return false
}
