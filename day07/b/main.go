// https://adventofcode.com/2022/day/7
// Day 7: No Space Left On Device
package main

import (
	"fmt"
	"oac/libaoc"
	"strings"
)

type File struct {
	name string
	size int
}
type Dir struct {
	name   string
	parent string
	files  []File
	dirs   []string
}

var dirs = make(map[string]Dir)

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	initialize(input)
	// fmt.Printf("%+v", dirs)
	// fmt.Printf("Sum of Dirs under and equal to 10k: %d\n", getSumDirsUnder100k())

	diskfree := 70000000 - getDirSize("//")
	needed := 30000000 - diskfree
	name, size := findSpace(needed)
	fmt.Printf("Diskfree: %d\nWe need another %d to have 30000000 free.\nDeleting %s would free %d\n", diskfree, needed, name, size)
}

func findSpace(needed int) (string, int) {
	best := 70000000
	var name string
	for dir := range dirs {
		size := getDirSize(dir)
		if size > needed {
			if size-needed < best-needed {
				best = size
				name = dir
				fmt.Printf("best: %d, %s\n", best, name)
			}
		}
	}
	return name, best
}

func getSumDirsUnder100k() (sum int) {
	for dir := range dirs {
		size := getDirSize(dir)
		if size <= 100000 {
			sum += size
		}
	}
	return sum
}

func getDirSize(dirname string) int {
	dir := dirs[dirname]
	var size int
	for _, file := range dir.files {
		size += file.size
	}
	for _, subdir := range dir.dirs {
		size += getDirSize(subdir)
	}
	return size
}

func initialize(input []string) {
	var active string
	dirs["//"] = Dir{name: "//"}
	for _, line := range input {
		if line[0:1] == "$" {
			switch {
			case line[2:4] == "cd": // Changing directory
				fields := strings.Fields(line)
				if fields[2] == ".." {
					active = dirs[active].parent
				} else {
					active += "/" + fields[2]
				}
			case line[2:4] == "ls": // Listing coming
				// We don't need this..
			}
		} else {
			fields := strings.Fields(line)
			if fields[0] == "dir" { // Dir found!
				// Add this dir to subdirs in active dir
				dir := dirs[active]
				dir.dirs = append(dir.dirs, active+"/"+fields[1])
				dirs[active] = dir

				// Add the active dir as parent to the subdir
				dir = dirs[active+"/"+fields[1]]
				dir.name = active + "/" + fields[1]
				dir.parent = active
				dirs[active+"/"+fields[1]] = dir
			} else { // File found!
				file := File{
					name: fields[1],
					size: libaoc.SilentAtoi(fields[0]),
				}
				dir := dirs[active]
				dir.files = append(dirs[active].files, file)
				dirs[active] = dir
			}
		}
	}
}
