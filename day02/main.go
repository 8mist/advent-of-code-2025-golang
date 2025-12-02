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

func part1(input string) int64 {
	ranges := parseInput(input)
	total := int64(0)

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		if len(parts) != 2 {
			continue
		}

		start, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}

		end, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}

		for i := start; i <= end; i++ {
			if isRepeatedTwice(strconv.Itoa(i)) {
				total += int64(i)
			}
		}
	}

	return total
}

func part2(input string) int64 {
	ranges := parseInput(input)

	total := int64(0)

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		if len(parts) != 2 {
			continue
		}

		start, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}

		end, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}

		for i := start; i <= end; i++ {
			if isRepeatedPattern(strconv.Itoa(i)) {
				total += int64(i)
			}
		}
	}

	return total
}

func parseInput(input string) []string {
	s := strings.ReplaceAll(input, "\r", "")
	s = strings.TrimSpace(s)
	return strings.Split(s, ",")
}

func isRepeatedTwice(nStr string) bool {
	n := len(nStr)
	if n%2 != 0 {
		return false
	}
	half := n / 2
	return nStr[:half] == nStr[half:]
}

func isRepeatedPattern(nStr string) bool {
	n := len(nStr)

	for seqLen := 1; seqLen*2 <= n; seqLen++ {
		if n%seqLen != 0 {
			continue
		}

		seq := nStr[:seqLen]
		isRepeat := true

		for i := seqLen; i < n; i += seqLen {
			if nStr[i:i+seqLen] != seq {
				isRepeat = false
				break
			}
		}

		if isRepeat {
			return true
		}
	}

	return false
}
