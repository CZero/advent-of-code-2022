// https://adventofcode.com/2022/day/8
// Day 8: Treetop Tree House
package main

import (
	"fmt"
	"oac/libaoc"
)

type Coord struct {
	c int // Column
	r int // Row
}

type Tree struct {
	height  int  // Tree height
	visable bool // Visable or not
	scenic  int
}

var (
	treeGrid   = make(map[Coord]Tree) // Map of trees, visible or no
	gridheight int
	gridwidth  int
)

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	gridheight = len(input)
	gridwidth = len(input[0])
	buildGrid(input)
	printGrid()
	fmt.Println(getVisableTrees())
	fmt.Println(getScenicScores())
}

func getScenicScores() (highestScenic int) {
	for r := 0; r < gridheight; r++ {
		for c := 0; c < gridwidth; c++ {
			score := calcScenicScores(Coord{c, r})
			tree := treeGrid[Coord{c, r}]
			tree.scenic = score
			treeGrid[Coord{c, r}] = tree
			if score > highestScenic {
				highestScenic = score
			}
		}
	}
	return highestScenic
}

func calcScenicScores(coord Coord) int {
	var top, bottom, left, right int // top, bottom, left, right
	// get up
	if coord.r != gridheight {
		for r := coord.r + 1; r < gridheight; r++ {
			top++
			if treeGrid[coord].height <= treeGrid[Coord{coord.c, r}].height {
				break
			}
		}
	}
	// fmt.Printf("To top: %d\n", top)

	// get down
	if coord.r != 0 {
		for r := coord.r - 1; r >= 0; r-- {
			bottom++
			// fmt.Printf("%v, %v\n", coord, Coord{coord.c, r})
			// fmt.Printf("Treeheight: %d, looking down: %d\n", treeGrid[coord].height, treeGrid[Coord{coord.c, r}].height)
			if treeGrid[coord].height <= treeGrid[Coord{coord.c, r}].height {
				break
			}
		}
	}
	// fmt.Printf("To bottom: %d\n", bottom)

	// get left
	if coord.c != 0 {
		for c := coord.c - 1; c >= 0; c-- {
			left++
			// fmt.Printf("%v, %v\n", coord, Coord{c, coord.r})
			// fmt.Printf("Treeheight: %d, looking left: %d\n", treeGrid[coord].height, treeGrid[Coord{c, coord.r}].height)
			if treeGrid[coord].height <= treeGrid[Coord{c, coord.r}].height {
				break
			}
		}
	}
	// fmt.Printf("To left: %d\n", left)

	// get right
	if coord.c != gridwidth {
		for c := coord.c + 1; c < gridwidth; c++ {
			right++
			// fmt.Printf("%v, %v\n", coord, Coord{c, coord.r})
			// fmt.Printf("Treeheight: %d, looking left: %d\n", treeGrid[coord].height, treeGrid[Coord{c, coord.r}].height)
			if treeGrid[coord].height <= treeGrid[Coord{c, coord.r}].height {
				break
			}
		}
	}
	// fmt.Printf("To right: %d\n", right)

	return top * bottom * left * right
}

func getVisableTrees() (num int) {
	for r := 0; r < gridheight; r++ {
		for c := 0; c < gridwidth; c++ {
			if isVisable(Coord{c, r}) {
				num++
				tree := treeGrid[Coord{c, r}]
				tree.visable = true
				treeGrid[Coord{c, r}] = tree
			}
		}
	}
	return num
}

func isVisable(coord Coord) bool {
	switch {
	case coord.c == 0 || coord.c == gridwidth-1:
		return true
	case coord.r == 0 || coord.r == gridheight-1:
		return true
	}
	// visible from the top?
	visible := true
	for r := coord.r + 1; r < gridheight; r++ {
		if treeGrid[Coord{coord.c, r}].height >= treeGrid[coord].height {
			visible = false
		}
	}
	if visible {
		return true
	}
	// visible from the bottom?
	visible = true
	for r := coord.r - 1; r >= 0; r-- {
		if treeGrid[Coord{coord.c, r}].height >= treeGrid[coord].height {
			visible = false
		}
	}
	if visible {
		return true
	}
	// visible from the left?
	visible = true
	for c := coord.c - 1; c >= 0; c-- {
		if treeGrid[Coord{c, coord.r}].height >= treeGrid[coord].height {
			visible = false
		}
	}
	if visible {
		return true
	}
	// visible from the right?
	visible = true
	for c := coord.c + 1; c <= gridwidth; c++ {
		if treeGrid[Coord{c, coord.r}].height >= treeGrid[coord].height {
			visible = false
		}
	}
	if visible {
		return true
	}
	return false
}

func buildGrid(input []string) {
	var row int
	for r := len(input) - 1; r >= 0; r-- {
		for c, heightstr := range input[r] {
			height := libaoc.SilentAtoi(string(heightstr))
			treeGrid[Coord{c, row}] = Tree{height: height}
		}
		row++
	}
}

func printGrid() {
	for r := gridheight - 1; r >= 0; r-- {
		var row []int
		for c := 0; c < gridwidth; c++ {
			row = append(row, treeGrid[Coord{c, r}].height)
		}
		fmt.Println(row)
	}
	fmt.Println()
}

func inGrid(coord Coord) bool {
	_, ok := treeGrid[coord]
	return ok
}
