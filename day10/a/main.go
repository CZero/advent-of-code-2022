// https://adventofcode.com/2022/day/10
// Day 10: Cathode-Ray Tube
package main

import (
	"aoc/libaoc"
	"fmt"
	"strings"
)

type Display map[int]int

var display = make(Display)

func main() {
	// input, err := libaoc.ReadLines("example1.txt")
	// input, err := libaoc.ReadLines("example2.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	runProgram(input)
	// display.print()
	display.signal()
	display.sumSignalStrength()
}

func runProgram(input []string) {
	// Init display
	tic := 1
	display[1] = 1

	// Go be a program
	for _, line := range input {
		if line == "noop" {
			tic = noop(tic)
			continue
		}
		instructions := strings.Fields(line)
		switch instructions[0] {
		case "addx":
			tic = addx(tic, libaoc.SilentAtoi(instructions[1]))
		default:
			panic("Something else than addx")
		}
	}
}

func noop(tic int) int {
	tic++
	display[tic] = display[tic-1]
	return tic
}

func addx(tic, num int) int {
	for i := 1; i < 2; i++ {
		tic++
		display[tic] = display[tic-1]
	}
	tic++
	display[tic] = display[tic-1] + num
	return tic
}

func (d Display) print() {
	for i := 0; i <= len(d); i++ {
		fmt.Printf("Line %d: %10d\n", i, d[i])
	}
}

func (d Display) signal() {
	for i := 20; i <= len(d); i += 40 {
		fmt.Printf("Line %d: %10d\n", i, d[i])
	}
}

func (d Display) sumSignalStrength() {
	var sum int
	for i := 20; i <= len(d); i += 40 {
		sum += i * d[i]
		fmt.Printf("%d * %10d = %d\n", i, d[i], i*d[i])

	}
	fmt.Printf("Signal strength sum: %d \n", sum)
}
