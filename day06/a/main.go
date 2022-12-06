// https://adventofcode.com/2022/day/
// Day x:
package main

import (
	"aoc/libaoc"
	"fmt"
)

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	fmt.Printf("Marker found at: %d\n", findStartMarkers(input[0]))
	// _,=findStartMarkers(input)
}

func findStartMarkers(input string) int {
	var (
		buffer []string
	)
	for i, char := range input {
		if len(buffer) < 4 {
			buffer = append(buffer, string(char))
			continue
		}
		if containsDoubles(buffer) {
			buffer = append(buffer, string(char))
			buffer = buffer[1:]
			continue
		}
		for _, buf := range buffer {
			if buf == string(char) {
				buffer = append(buffer, string(char))
				buffer = buffer[1:]
				continue
			}
		}
		return i
	}
	return 0
}

func containsDoubles(input []string) bool {
	for i, char := range input {
		for j, comp := range input {
			if i != j {
				if char == comp {
					return true
				}
			}
		}
	}
	return false
}
