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

const (
	YOU = "you"
	OUT = "out"
	SVR = "svr"
	DAC = "dac"
	FFT = "fft"
)

type Devices map[string][]string
type Cache1 map[string]int

type Key2 struct {
	node    string
	seenDAC bool
	seenFFT bool
}
type Cache2 map[Key2]int

func part1(input string) int {
	devices := parseInput(input)
	cache := make(Cache1)
	return countPathsFrom(devices, YOU, cache)
}

func part2(input string) int {
	devices := parseInput(input)
	cache := make(Cache2)
	return countPathsWithRequired(devices, SVR, false, false, cache)
}

func parseInput(input string) Devices {
	s := strings.ReplaceAll(input, "\r", "")
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")

	devices := make(Devices)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			continue
		}
		device := strings.TrimSpace(parts[0])
		outputs := strings.Fields(parts[1])
		devices[device] = outputs
	}
	return devices
}

func countPathsFrom(devices Devices, current string, cache Cache1) int {
	if current == OUT {
		return 1
	}
	if val, ok := cache[current]; ok {
		return val
	}
	outputs, ok := devices[current]
	if !ok {
		return 0
	}
	total := 0
	for _, next := range outputs {
		total += countPathsFrom(devices, next, cache)
	}
	cache[current] = total
	return total
}

func countPathsWithRequired(devices Devices, current string, seenDAC, seenFFT bool, memo Cache2) int {
	if current == DAC {
		seenDAC = true
	}
	if current == FFT {
		seenFFT = true
	}

	if current == OUT {
		if seenDAC && seenFFT {
			return 1
		}
		return 0
	}

	key := Key2{current, seenDAC, seenFFT}
	if v, ok := memo[key]; ok {
		return v
	}

	children, ok := devices[current]
	if !ok || len(children) == 0 {
		return 0
	}

	total := 0
	for _, next := range children {
		total += countPathsWithRequired(devices, next, seenDAC, seenFFT, memo)
	}

	memo[key] = total
	return total
}
