package main

import (
	_ "embed"
	"flag"
	"fmt"
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

func part1(input string) int64 {
	lines := parseInput(input)

	total := int64(0)
	for _, line := range lines {
		total += findJoltage(strings.TrimSpace(line), 2)
	}
	return total
}

func part2(input string) int64 {
	lines := parseInput(input)

	total := int64(0)
	for _, line := range lines {
		total += findJoltage(strings.TrimSpace(line), 12)
	}
	return total
}

func parseInput(input string) []string {
	s := strings.ReplaceAll(input, "\r", "")
	s = strings.TrimSpace(s)
	return strings.Split(s, "\n")
}

func findJoltage(line string, numValues int) int64 {
	result := int64(0)
	remaining := numValues
	start := 0

	for remaining > 0 {
		maxDigit := byte('0')
		maxIndex := start
		endSearch := len(line) - remaining + 1

		for i := start; i < endSearch; i++ {
			if line[i] > maxDigit {
				maxDigit = line[i]
				maxIndex = i
			}
		}

		result = result*10 + int64(maxDigit-'0')
		start = maxIndex + 1
		remaining--
	}

	return result
}
