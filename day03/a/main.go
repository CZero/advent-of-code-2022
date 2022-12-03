// https://adventofcode.com/2022/day/3
// Day 3: Rucksack Reorganization
package main

import (
	"aoc/libaoc"
	"fmt"
)

type Rucksack map[string]bool

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	misplaced := getMisplacedItems(input)
	fmt.Printf("Itemsprios: %v\nTotal prio: %d\n", misplaced, sumItems(misplaced))
}

func getMisplacedItems(input []string) (misplaced []int) {
	for _, rucksack := range input {
		firstHalf, secondHalf := cutInHalf(rucksack)
		var compartmentA = make(Rucksack) // We will check against this one
		var foundItems = make(Rucksack)   // Here we put the items we found ONCE, not every occurance
		for _, item := range firstHalf {
			compartmentA.add(string(item))
		}
		for _, item := range secondHalf {
			if compartmentA.has(string(item)) {
				if !foundItems.has(string(item)) {
					misplaced = append(misplaced, value(string(item)))
					foundItems.add(string(item))
				}
			}
		}
	}
	return misplaced
}

func cutInHalf(input string) (a, b string) {
	half := len(input) / 2
	a = input[:half]
	b = input[half:]
	return a, b
}

func value(input string) int {
	for _, charNumber := range input {
		if charNumber < 91 { // lowercase
			return int(charNumber) - 38
		} else { // uppercase
			return int(charNumber) - 96
		}
	}
	panic("Shouldn't reach this")
}

func sumItems(input []int) (sum int) {
	for _, item := range input {
		sum += item
	}
	return sum
}

func (r Rucksack) has(item string) bool {
	_, ok := r[item]
	return ok
}

func (r Rucksack) add(item string) {
	r[item] = true
}
