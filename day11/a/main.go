// https://adventofcode.com/2022/day/11
// Day 11: Monkey in the Middle
package main

import (
	"aoc/libaoc"
	"fmt"
	"sort"
	"strings"
)

type Monkey struct {
	items       []int // Items (worrylevels) in hand
	optimes     bool  // Times or plus?
	opnum       int   // Operate by what? (-1 means old value, so itself)
	divider     int   // Divides by to decide to throw
	truefalse   []int // Throws to truefalse[0] on true, to truefalse[1] on false
	inspections int   // Number of times inspected
}

var (
	monkeys []Monkey
)

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	initMonkeys(input)

	runProgram(20)
	getMonkeyBusiness()
}

// getMonkeyBusiness get's the highest two inspections and multiplies them
func getMonkeyBusiness() {
	var inspections []int
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.inspections)
	}
	sort.Ints(inspections)
	fmt.Printf("The level of monkeybusiness: %d * %d = %d\n", inspections[len(inspections)-2], inspections[len(inspections)-1], inspections[len(inspections)-2]*inspections[len(inspections)-1])
}

// runProgram Runs the program for the input number of rounds
func runProgram(rounds int) {
	for i := 1; i <= rounds; i++ {
		doRound()

		// fmt.Printf("Ronde %d gedaan:\n", i)
		// printMonkeys()
		// println()
	}
}

// doRound does all monkeys in order once
func doRound() {
	for i := 0; i < len(monkeys); i++ {
		// fmt.Printf("doMonkey: %d\n", i)
		doMonkey(i)
		// fmt.Printf("Monkey %d gedaan:\n", i)
		// printMonkeys()
		// println()
	}
}

// doMonkey does the inspection and handling of all items the monkey has and updates all monkeys accordingly
func doMonkey(n int) {
	monkey := monkeys[n] // Take the monkey from the pile
	for _, item := range monkey.items {
		// first the operation
		if monkey.optimes { // Times
			if monkey.opnum == -1 { // By itself
				item *= item
			} else { // By the number
				item *= monkey.opnum
			}
		} else { // Plus
			if monkey.opnum == -1 { // By itself
				item += item
			} else { // By the number
				item += monkey.opnum
			}
		}

		// second the division by 3 and rounding down
		item /= 3

		// division check, add to monkey true or to monkey false's stash (in the end)
		if item%monkey.divider == 0 { // To monkey true
			monkeys[monkey.truefalse[0]].items = append(monkeys[monkey.truefalse[0]].items, item)
		} else { // To monkey false
			monkeys[monkey.truefalse[1]].items = append(monkeys[monkey.truefalse[1]].items, item)
		}

		// and finally: up the inspection for this item
		monkey.inspections++
	}
	// The monkey threw away all his items
	monkey.items = []int{}
	monkeys[n] = monkey // Add the monkey back to the pile
}

// initMonkeys initializes the monkeys according to the input passed
func initMonkeys(input []string) {
	var monkey Monkey
	i := 1
	for _, line := range input {
		switch i {
		case 2: // Items
			// fmt.Println("items: " + line)
			fields := strings.Fields(line)
			for j, item := range fields[2:] {
				if j != len(fields[2:])-1 { // Not last, has a trailing ","
					monkey.items = append(monkey.items, libaoc.SilentAtoi(item[:len(item)-1]))
				} else { // Last, you can just add this one
					monkey.items = append(monkey.items, libaoc.SilentAtoi(item))
				}
			}
		case 3: // Operation
			fields := strings.Fields(line)
			if fields[4] == "*" {
				monkey.optimes = true
			}
			if fields[5] == "old" {
				monkey.opnum = -1
			} else {
				monkey.opnum = libaoc.SilentAtoi(fields[5])
			}
		case 4: // Test
			fields := strings.Fields(line)
			monkey.divider = libaoc.SilentAtoi(fields[len(fields)-1])
		case 5: // Monkey if true
			fields := strings.Fields(line)
			monkey.truefalse = append(monkey.truefalse, libaoc.SilentAtoi(fields[len(fields)-1]))
		case 6: // Monkey if false
			fields := strings.Fields(line)
			monkey.truefalse = append(monkey.truefalse, libaoc.SilentAtoi(fields[len(fields)-1]))

			// And finish the monkey
			fmt.Printf("This monkey is complete:\n%+v\n\n", monkey)
			monkeys = append(monkeys, monkey)
			monkey = Monkey{}
			i = -1 // restart the count, passing an empty line before reaching 1 again
		}
		i++
	}
}

func printMonkeys() {
	for i, monkey := range monkeys {
		fmt.Printf("Monkey %d: %#v\n", i, monkey)
	}
}
