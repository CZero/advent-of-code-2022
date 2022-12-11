// https://adventofcode.com/2022/day/9
// Day 9: Rope Bridge
package main

import (
	"aoc/libaoc"
	"fmt"
	"math"
	"strings"
)

type Coord struct {
	c int // Column
	r int // Row
}

type Visitmap map[Coord]bool

var tailVisited = make(Visitmap)

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	doMotions(input)
}

func doMotions(input []string) {
	var (
		head Coord
		tail Coord
	)
	tailVisited[Coord{0, 0}] = true
	for _, line := range input {
		head, tail = moveHead(line, head, tail)
	}
	fmt.Println(tailVisited.length())
	// tailVisited.print()
}

func moveTail(head, tail Coord) Coord {
	if (math.Abs(float64(head.c-tail.c)) >= 1 && math.Abs(float64(head.r-tail.r)) > 1) || math.Abs(float64(head.c-tail.c)) > 1 && math.Abs(float64(head.r-tail.r)) >= 1 { // Math.Abs converts a negative into a positive number.
		// TODO Hier moet de schuin magie komen.
		// The x-dif AND y-dif is to big
		switch {
		case head.c > tail.c:
			tail.c++
		case head.c < tail.c:
			tail.c--
		}
		switch {
		case head.r > tail.r:
			tail.r++
		case head.r < tail.r:
			tail.r--
		}
		tailVisited[tail] = true
		return tail
	}

	if math.Abs(float64(head.c-tail.c)) > 1 { // Math.Abs converts a negative into a positive number.
		// The x-dif is to big
		switch {
		case head.c > tail.c:
			tail.c++
			tailVisited[tail] = true
		case head.c < tail.c:
			tail.c--
			tailVisited[tail] = true
		}
		return tail
	}
	if math.Abs(float64(head.r-tail.r)) > 1 {
		// The y-dif is to big
		switch {
		case head.r > tail.r:
			tail.r++
			tailVisited[tail] = true
		case head.r < tail.r:
			tail.r--
			tailVisited[tail] = true
		}
		return tail
	}
	return tail
}

func moveHead(input string, head, tail Coord) (Coord, Coord) {
	instructions := strings.Fields(input)
	switch instructions[0] {
	case "U":
		for i := 0; i < libaoc.SilentAtoi(instructions[1]); i++ {
			head = Coord{head.c, head.r + 1}
			tail = moveTail(head, tail)
			fmt.Printf("Head: %v Tail: %v\n", head, tail)
		}
		return head, tail
	case "D":
		for i := 0; i < libaoc.SilentAtoi(instructions[1]); i++ {
			head = Coord{head.c, head.r - 1}
			tail = moveTail(head, tail)
			fmt.Printf("Head: %v Tail: %v\n", head, tail)
		}
		return head, tail
	case "L":
		for i := 0; i < libaoc.SilentAtoi(instructions[1]); i++ {
			head = Coord{head.c - 1, head.r}
			tail = moveTail(head, tail)
			fmt.Printf("Head: %v Tail: %v\n", head, tail)
		}
		return head, tail
	case "R":
		for i := 0; i < libaoc.SilentAtoi(instructions[1]); i++ {
			head = Coord{head.c + 1, head.r}
			tail = moveTail(head, tail)
			fmt.Printf("Head: %v Tail: %v\n", head, tail)
		}
		return head, tail
	default:
		panic("There should be no default! " + instructions[0])
	}
}

// func (v Visitmap) has(visted coord) bool {
// 	_, ok := v[visted]
// 	return ok
// }

func (v Visitmap) length() int {
	return len(v)
}
func (v Visitmap) print() {
	for visited, _ := range v {
		fmt.Printf("Tailvisits: %v\n", visited)
	}
}
