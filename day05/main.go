package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
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

type Interval struct {
	From int64
	To   int64
}

func part1(input string) int64 {
	ranges, ids := parseInput(input)

	var total int64
	for _, id := range ids {
		if isFresh(id, ranges) {
			total++
		}
	}
	return total
}

func part2(input string) int64 {
	ranges, _ := parseInput(input)
	merged := mergeIntervals(ranges)

	var total int64
	for _, r := range merged {
		total += r.To - r.From + 1
	}

	return total
}

func parseInput(input string) ([]Interval, []int64) {
	blocks := strings.Split(strings.TrimSpace(strings.ReplaceAll(input, "\r", "")), "\n\n")
	if len(blocks) != 2 {
		panic("expected two blocks separated by a blank line")
	}

	ranges := parseBlockRanges(blocks[0])
	ids := parseBlockIDs(blocks[1])
	return ranges, ids
}

func parseBlockRanges(block string) []Interval {
	var ranges []Interval

	for _, line := range strings.Split(block, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			panic("invalid range line: " + line)
		}

		from, err1 := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
		to, err2 := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
		if err1 != nil || err2 != nil {
			panic("invalid range numbers: " + line)
		}

		ranges = append(ranges, Interval{From: from, To: to})
	}

	return ranges
}

func parseBlockIDs(block string) []int64 {
	var ids []int64

	for _, line := range strings.Split(block, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		id, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic("invalid id line: " + line)
		}

		ids = append(ids, id)
	}

	return ids
}

func isFresh(id int64, ranges []Interval) bool {
	for _, r := range ranges {
		if id >= r.From && id <= r.To {
			return true
		}
	}
	return false
}

func mergeIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].From < intervals[j].From
	})

	var merged []Interval
	current := intervals[0]

	for i := 1; i < len(intervals); i++ {
		next := intervals[i]
		if current.To >= next.From {
			if current.To < next.To {
				current.To = next.To
			}
		} else {
			merged = append(merged, current)
			current = next
		}
	}
	merged = append(merged, current)

	return merged
}
