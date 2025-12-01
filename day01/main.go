package main

import (
	"8mist/aoc/mathy"
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var reInstruction = regexp.MustCompile(`([LR])(\d+)`)

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
	instructions := parseInput(input)

	position := 50
	zeroCount := 0

	for _, inst := range instructions {
		dir, dist := parseInstruction(inst)
		delta := directionDelta(dir)

		position = mathy.Mod(position+delta*dist, 100)
		if position == 0 {
			zeroCount++
		}
	}

	return zeroCount
}

func part2(input string) int {
	instructions := parseInput(input)

	position := 50
	zeroCount := 0

	for _, inst := range instructions {
		dir, dist := parseInstruction(inst)
		delta := directionDelta(dir)

		for i := 1; i <= dist; i++ {
			intermediate := mathy.Mod(position+delta*i, 100)
			if intermediate == 0 {
				zeroCount++
			}
		}

		position = mathy.Mod(position+delta*dist, 100)
	}

	return zeroCount
}

func parseInput(input string) []string {
	s := strings.ReplaceAll(input, "\r", "")
	s = strings.TrimSpace(s)
	return strings.Split(s, "\n")
}

func parseInstruction(line string) (direction string, distance int) {
	m := reInstruction.FindStringSubmatch(line)
	if m == nil {
		return "", 0
	}
	dist, _ := strconv.Atoi(m[2])
	return m[1], dist
}

func directionDelta(dir string) int {
	if dir == "L" {
		return -1
	}
	return 1
}
