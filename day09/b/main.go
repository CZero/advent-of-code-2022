// https://adventofcode.com/2022/day/9
// Day 9: Rope Bridge
package main

import (
	"fmt"
	"math"
	"oac/libaoc"
	"strings"
)

type Coord struct {
	c int // Column
	r int // Row
}

type Visitmap map[Coord]bool

var (
	tailVisited = make(Visitmap)
	knots       = map[string]Coord{
		"1": {},
		"2": {},
		"3": {},
		"4": {},
		"5": {},
		"6": {},
		"7": {},
		"8": {},
		"9": {},
	}
	knotslist = []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
	}
)

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	// input, err := libaoc.ReadLines("example2.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	doMotions(input)
}

func doMotions(input []string) {
	var (
		head Coord
	)
	tailVisited[Coord{0, 0}] = true

	for _, line := range input {
		head = moveHead(line, head)
	}
	fmt.Println(tailVisited.length())
	// tailVisited.print()
}

func moveKnots(head Coord) {
	for i := 0; i < 9; i++ {
		if i == 0 {
			knots[knotslist[i]] = moveTail(head, knots[knotslist[i]], false)
		} else if i == 8 { // If the last one moves the tail should be logged, toggle that
			knots[knotslist[i]] = moveTail(knots[knotslist[i-1]], knots[knotslist[i]], true)
		} else {
			knots[knotslist[i]] = moveTail(knots[knotslist[i-1]], knots[knotslist[i]], false)
		}
	}
}

func moveTail(head, tail Coord, realdeal bool) Coord {
	if (math.Abs(float64(head.c-tail.c)) >= 1 && math.Abs(float64(head.r-tail.r)) > 1) || math.Abs(float64(head.c-tail.c)) > 1 && math.Abs(float64(head.r-tail.r)) >= 1 { // Math.Abs converts a negative into a positive number.
		// The weird diagonal magic rule
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
		if realdeal {
			tailVisited[tail] = true
		}
		return tail
	}

	if math.Abs(float64(head.c-tail.c)) > 1 { // Math.Abs converts a negative into a positive number.
		// The x-dif is to big
		switch {
		case head.c > tail.c:
			tail.c++
			if realdeal {
				tailVisited[tail] = true
			}
		case head.c < tail.c:
			tail.c--
			if realdeal {
				tailVisited[tail] = true
			}
		}
		return tail
	}
	if math.Abs(float64(head.r-tail.r)) > 1 {
		// The y-dif is to big
		switch {
		case head.r > tail.r:
			tail.r++
			if realdeal {
				tailVisited[tail] = true
			}
		case head.r < tail.r:
			tail.r--
			if realdeal {
				tailVisited[tail] = true
			}
		}
		return tail
	}
	return tail
}

func moveHead(input string, head Coord) Coord {
	instructions := strings.Fields(input)
	switch instructions[0] {
	case "U":
		for i := 0; i < libaoc.SilentAtoi(instructions[1]); i++ {
			head = Coord{head.c, head.r + 1}
			moveKnots(head)
			fmt.Printf("Head: %v\n", head)
			for _, knot := range knotslist {
				fmt.Printf("Knot %s: %v\n", knot, knots[knot])
			}
			fmt.Println()
		}
		return head
	case "D":
		for i := 0; i < libaoc.SilentAtoi(instructions[1]); i++ {
			head = Coord{head.c, head.r - 1}
			moveKnots(head)
			fmt.Printf("Head: %v\n", head)
			for _, knot := range knotslist {
				fmt.Printf("Knot %s: %v\n", knot, knots[knot])
			}
			fmt.Println()
		}
		return head
	case "L":
		for i := 0; i < libaoc.SilentAtoi(instructions[1]); i++ {
			head = Coord{head.c - 1, head.r}
			moveKnots(head)
			fmt.Printf("Head: %v\n", head)
			for _, knot := range knotslist {
				fmt.Printf("Knot %s: %v\n", knot, knots[knot])
			}
			fmt.Println()
		}
		return head
	case "R":
		for i := 0; i < libaoc.SilentAtoi(instructions[1]); i++ {
			head = Coord{head.c + 1, head.r}
			moveKnots(head)
			fmt.Printf("Head: %v\n", head)
			for _, knot := range knotslist {
				fmt.Printf("Knot %s: %v\n", knot, knots[knot])
			}
			fmt.Println()
		}
		return head
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
