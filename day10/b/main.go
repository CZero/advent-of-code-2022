// https://adventofcode.com/2022/day/10
// Day 10: Cathode-Ray Tube
package main

import (
	"fmt"
	"oac/libaoc"
	"strings"
)

type Display map[int]int
type coord struct {
	c int // Column
	r int // Row
}
type CRT map[coord]string

var display = make(Display)
var crt = make(CRT)

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
	crt.drawScreen()
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
		fmt.Printf("%d * %d = %d\n", i, d[i], i*d[i])

	}
	fmt.Printf("Signal strength sum: %d \n", sum)
}

func (c CRT) addPixel(cycle int) {
	row := (cycle - 1) / 40
	col := (cycle - 1) % 40
	if col == display[cycle] || col == display[cycle]-1 || col == display[cycle]+1 {
		c[coord{col, row}] = "#"
		return
	}
	// fmt.Println(cycle, row, col) // Tested and leads to a fully drawn roster
	crt[coord{col, row}] = "."
}

func (c CRT) drawScreen() {
	for i := 1; i <= 240; i++ {
		c.addPixel(i)
	}
	// fmt.Printf("%+v", c)
	// Drawing time!
	fmt.Println()
	for row := 0; row < 6; row++ {
		for col := 0; col < 40; col++ {
			fmt.Printf("%s", c[coord{col, row}])
			if col == 39 {
				fmt.Printf("\n")
			}
		}
	}
	fmt.Println()
}
