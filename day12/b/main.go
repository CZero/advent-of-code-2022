// https://adventofcode.com/2022/day/12
// Day 12: Hill Climbing Algorithm
package main

import (
	"fmt"
	"oac/libaoc"
)

// Coordinates
type Coord struct {
	c int // column
	r int // row
}

// Locations
type Loc struct {
	height   string
	distance int
	options  []Coord
}

type Hill map[Coord]Loc

// Path
type Path []Coord

var (
	start      Coord
	end        Coord
	height     int
	width      int
	pathsFound []Path
	hill       = make(Hill)
)

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	buildHill(input)
	// fmt.Printf("%+v", hill)
	fmt.Printf("Start: %+v\n", hill[start])
	fmt.Printf("End: %+v\n", hill[end])
	hill.draw()
	As := findAs()
	// fmt.Println(As)
	doAllStarts(As)
	// findEnd(path)
	// fmt.Printf("%v\n", hill[end])
	// fmt.Printf("The shortest path = %d\n", hill[end].distance-1)
	// fmt.Printf(path)
	// hill.drawDistances()
}

func doAllStarts(As []Coord) {
	var shortest int
	for _, start := range As {
		resetHill()
		startLoc := hill[start]
		startLoc.distance = 1
		hill[start] = startLoc
		path := Path{start}
		distanceFinder(path)
		if hill[end].distance == 0 {
			continue
		}
		fmt.Printf("The shortest path from %v= %d\n", start, hill[end].distance-1)
		if shortest == 0 || shortest > hill[end].distance-1 {
			shortest = hill[end].distance - 1
		}
	}
	fmt.Printf("The shortest was: %d\n", shortest)
}

func resetHill() {
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			hillLoc := hill[Coord{c, r}]
			hillLoc.distance = 0
			hill[Coord{c, r}] = hillLoc
		}
	}
}

func findAs() (As []Coord) {
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if hill[Coord{c, r}].height == "a" {
				As = append(As, Coord{c, r})
			}
		}
	}
	return As
}

func distanceFinder(path Path) {
	for _, option := range hill[path[len(path)-1]].options {
		switch {
		case hill[option].distance == 0: // The location is unvisited
			optionLoc := hill[option]
			optionLoc.distance = hill[path[len(path)-1]].distance + 1
			hill[option] = optionLoc
			further := append(path, option)
			distanceFinder(further)
		case hill[option].distance > hill[path[len(path)-1]].distance+1:
			optionLoc := hill[option]
			optionLoc.distance = hill[path[len(path)-1]].distance + 1
			hill[option] = optionLoc
			further := append(path, option)
			distanceFinder(further)
		default:
			// The next step is at a lower distance than we are. This route can die.
		}
	}
}

func buildHill(input []string) {
	height = len(input)
	width = len(input[0])
	r := 0
	for i := height - 1; i >= 0; i-- {
		for c, char := range input[i] {
			if string(char) == "S" {
				char = 'a'
			}
			if string(char) == "E" {
				end = Coord{c, r}
				char = 'z'
			}
			hill[Coord{c, r}] = Loc{height: string(char)}
		}
		r++
	}
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			loc := hill[Coord{c, r}]
			loc.options = getOptions(Coord{c, r})
			hill[Coord{c, r}] = loc
		}
	}
}

func getOptions(loc Coord) (options []Coord) {
	locHeight := hill[loc].height
	if loc.r < height-1 { // Look above
		if getLetterValue(locHeight) > getLetterValue(hill[Coord{loc.c, loc.r + 1}].height) || getLetterValue(hill[Coord{loc.c, loc.r + 1}].height)-getLetterValue(locHeight) <= 1 {
			options = append(options, Coord{loc.c, loc.r + 1})
		}
	}
	if loc.r > 0 { // Look below
		if getLetterValue(locHeight) > getLetterValue(hill[Coord{loc.c, loc.r - 1}].height) || getLetterValue(hill[Coord{loc.c, loc.r - 1}].height)-getLetterValue(locHeight) <= 1 {
			options = append(options, Coord{loc.c, loc.r - 1})
		}
	}
	if loc.c < width-1 { // Look right
		if getLetterValue(locHeight) > getLetterValue(hill[Coord{loc.c + 1, loc.r}].height) || getLetterValue(hill[Coord{loc.c + 1, loc.r}].height)-getLetterValue(locHeight) <= 1 {
			options = append(options, Coord{loc.c + 1, loc.r})
		}
	}
	if loc.c > 0 { // Look left
		if getLetterValue(locHeight) > getLetterValue(hill[Coord{loc.c - 1, loc.r}].height) || getLetterValue(hill[Coord{loc.c - 1, loc.r}].height)-getLetterValue(locHeight) <= 1 {
			options = append(options, Coord{loc.c - 1, loc.r})
		}
	}
	return options
}

func (p Path) visited(coord Coord) bool {
	for _, step := range p {
		if coord == step {
			// fmt.Println("Hier zijn we al geweest")
			return true
		}
	}
	return false
}

func getLetterValue(s string) int {
	for _, char := range s {
		return int(char)
	}
	return 0
}

func (h Hill) draw() {
	for r := height; r >= 0; r-- {
		for c := 0; c <= width; c++ {
			fmt.Printf("%s", h[Coord{c, r}].height)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\nStart: %v\nEnd: %v\nHeight: %d, Width: %d\n", start, end, height, width)
}
func (h Hill) drawDistances() {
	for r := height; r >= 0; r-- {
		for c := 0; c <= width; c++ {
			fmt.Printf("%3d", h[Coord{c, r}].distance)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\nStart: %v\nEnd: %v\nHeight: %d, Width: %d\n", start, end, height, width)
}
