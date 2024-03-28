package main

import (
	"fmt"
	"strings"
)

func getInput() string {
	return `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
}

type Thing = int

const (
	Tree Thing = iota
	Snow
)

func main() {
	treeCount := 0

	for r, line := range strings.Split(getInput(), "\n") {
		colLen := len(line)
		if string(line[r*3%colLen]) == "#" {
			treeCount++
		}

	}

	fmt.Printf("tree count %v\n", treeCount)
}
