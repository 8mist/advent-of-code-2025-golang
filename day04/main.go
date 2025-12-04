package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

type Grid [][]byte

//go:embed input.txt
var input string

var dirs = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

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

func part1(input string) int64 {
	lines := parseInput(input)
	grid := makeGrid(lines)
	return simulateRemovals(grid, 1)
}

func part2(input string) int64 {
	lines := parseInput(input)
	grid := makeGrid(lines)
	return simulateRemovals(grid, -1)
}

func parseInput(input string) []string {
	s := strings.ReplaceAll(input, "\r", "")
	s = strings.TrimSpace(s)
	return strings.Split(s, "\n")
}

func makeGrid(lines []string) Grid {
	h := len(lines)
	if h == 0 {
		return nil
	}

	grid := make([][]byte, h)
	for i, line := range lines {
		grid[i] = []byte(line)
	}

	return grid
}

func gridDims(g Grid) (int, int) {
	if len(g) == 0 {
		return 0, 0
	}
	return len(g), len(g[0])
}

func simulateRemovals(grid Grid, maxSteps int) int64 {
	total := int64(0)
	step := 0

	for {
		step++
		if maxSteps > 0 && step > maxSteps {
			break
		}

		toRemove := findAccessible(grid)
		if len(toRemove) == 0 {
			break
		}

		removeRolls(grid, toRemove)
		total += int64(len(toRemove))
	}

	return total
}

func findAccessible(grid [][]byte) [][2]int {
	h, w := gridDims(grid)
	toRemove := make([][2]int, 0, h*w/10)

	for y := range h {
		for x := range w {
			if grid[y][x] != '@' {
				continue
			}

			if countNeighbors(grid, h, w, y, x) < 4 {
				toRemove = append(toRemove, [2]int{y, x})
			}
		}
	}

	return toRemove
}

func countNeighbors(grid Grid, h int, w int, y, x int) int {
	count := 0

	for _, d := range dirs {
		ny, nx := y+d[0], x+d[1]
		if ny >= 0 && ny < h && nx >= 0 && nx < w && grid[ny][nx] == '@' {
			count++
		}
	}

	return count
}

func removeRolls(grid [][]byte, positions [][2]int) {
	for _, pos := range positions {
		grid[pos[0]][pos[1]] = '.'
	}
}
