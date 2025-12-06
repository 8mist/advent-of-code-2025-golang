package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

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
	lines := parseInput(input)

	var grid [][]string
	for _, line := range lines {
		grid = append(grid, strings.Fields(line))
	}

	numbers := grid[:len(grid)-1]
	ops := grid[len(grid)-1]

	return sumReduce(ops, numbers)
}

func part2(input string) int {
	lines := parseInput(input)

	maxLen := 0
	for _, line := range lines {
		maxLen = max(maxLen, len(line))
	}

	tally := make([]int, 0)
	total := 0

	for col := maxLen - 1; col >= 0; col-- {
		var column []string
		allSpaces := true
		for _, line := range lines {
			if col < len(line) {
				c := string(line[col])
				column = append(column, c)
				if c != " " {
					allSpaces = false
				}
			}
		}

		if allSpaces {
			continue
		}

		opPos := -1
		for i, c := range column {
			if c == "+" || c == "*" {
				opPos = i
				break
			}
		}

		numStr := ""
		for i := 0; i < len(column); i++ {
			if i == opPos {
				break
			}
			c := column[i]
			if c >= "0" && c <= "9" {
				numStr += c
			}
		}

		if numStr != "" {
			num, _ := strconv.Atoi(numStr)
			tally = append(tally, num)
		}

		if opPos != -1 {
			op := column[opPos]
			if len(tally) > 0 {
				f := getOp(op)
				result := reduceInts(tally, f)
				total += result
				tally = tally[:0]
			}
		}
	}

	return total
}

func parseInput(input string) []string {
	s := strings.ReplaceAll(input, "\r", "")
	return strings.Split(s, "\n")
}

func getOp(op string) func(int, int) int {
	switch op {
	case "+":
		return func(a, b int) int { return a + b }
	case "*":
		return func(a, b int) int { return a * b }
	default:
		panic("unknown operator: " + op)
	}
}

func sumReduce(ops []string, numbers [][]string) int {
	total := 0
	for i, op := range ops {
		f := getOp(op)
		col := column(numbers, i)
		acc := reduceInts(col, f)
		total += acc
	}
	return total
}

func column(matrix [][]string, idx int) []int {
	var col []int
	for _, row := range matrix {
		if idx < len(row) {
			val, _ := strconv.Atoi(row[idx])
			col = append(col, val)
		}
	}
	return col
}

func reduceInts(nums []int, f func(int, int) int) int {
	if len(nums) == 0 {
		return 0
	}
	acc := nums[0]
	for _, n := range nums[1:] {
		acc = f(acc, n)
	}
	return acc
}
