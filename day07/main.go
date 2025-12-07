package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

const (
	START    = "S"
	SPLITTER = "^"
)

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		fmt.Println("Output:", part1(input))
	} else {
		fmt.Println("Output:", part2(input))
	}
}

func part1(input string) int {
	grid := parseInput(input)
	height := len(grid)
	width := len(grid[0])
	startCol := findStartCol(grid, width)
	activeBeams := map[int]struct{}{startCol: {}}
	splits := 0

	for row := 1; row < height; row++ {
		newBeams := make(map[int]struct{})

		for col := range activeBeams {
			if col < 0 || col >= width {
				continue
			}

			if grid[row][col] != SPLITTER {
				newBeams[col] = struct{}{}
				continue
			}

			splits++

			if col-1 >= 0 {
				newBeams[col-1] = struct{}{}
			}

			if col+1 < width {
				newBeams[col+1] = struct{}{}
			}
		}

		activeBeams = newBeams
	}

	return splits
}

func part2(input string) int {
	grid := parseInput(input)
	height := len(grid)
	width := len(grid[0])
	startCol := findStartCol(grid, width)
	activeBeams := map[int]int{startCol: 1}

	for row := 1; row < height; row++ {
		newBeams := make(map[int]int)

		for col, count := range activeBeams {
			if grid[row][col] != SPLITTER {
				newBeams[col] += count
				continue
			}

			if col-1 >= 0 {
				newBeams[col-1] += count
			}

			if col+1 < width {
				newBeams[col+1] += count
			}
		}

		activeBeams = newBeams
	}

	totalTimelines := 0
	for _, count := range activeBeams {
		totalTimelines += count
	}

	return totalTimelines
}

func parseInput(input string) [][]string {
	s := strings.ReplaceAll(input, "\r", "")
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")

	var grid [][]string
	for _, line := range lines {
		line := strings.TrimSpace(line)
		if line == "" {
			continue
		}
		cells := strings.Split(line, "")
		grid = append(grid, cells)
	}
	return grid
}

func findStartCol(grid [][]string, width int) int {
	startCol := -1
	for col := range width {
		if grid[0][col] == START {
			startCol = col
			break
		}
	}
	return startCol
}
