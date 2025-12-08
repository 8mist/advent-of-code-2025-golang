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
		fmt.Println("Output:", part1(input, 1000))
	} else {
		fmt.Println("Output:", part2(input, 0))
	}
}

type point struct {
	x, y, z int
}

type pair struct {
	d2   int64
	a, b int
}

func part1(input string, k int) int {
	points := parsePoints(input)
	n := len(points)
	pairs := make([]pair, 0, n*(n-1)/2)

	for i := range n {
		for j := i + 1; j < n; j++ {
			d2 := squaredDistance(points[i], points[j])
			pairs = append(pairs, pair{d2: d2, a: i, b: j})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].d2 < pairs[j].d2
	})

	parent := make([]int, n)
	size := make([]int, n)
	for i := range n {
		parent[i] = i
		size[i] = 1
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		ra := find(a)
		rb := find(b)
		if ra == rb {
			return
		}
		if size[ra] < size[rb] {
			ra, rb = rb, ra
		}
		parent[rb] = ra
		size[ra] += size[rb]
	}

	limit := min(k, len(pairs))
	for i := range limit {
		union(pairs[i].a, pairs[i].b)
	}

	compSize := make(map[int]int)
	for i := range n {
		r := find(i)
		compSize[r]++
	}

	sizes := make([]int, 0, len(compSize))
	for _, s := range compSize {
		sizes = append(sizes, s)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	result := 1
	for i := 0; i < 3 && i < len(sizes); i++ {
		result *= sizes[i]
	}

	return result
}

func part2(input string, _ int) int {
	points := parsePoints(input)
	n := len(points)
	pairs := make([]pair, 0, n*(n-1)/2)

	for i := range n {
		for j := i + 1; j < n; j++ {
			d2 := squaredDistance(points[i], points[j])
			pairs = append(pairs, pair{d2: d2, a: i, b: j})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].d2 < pairs[j].d2
	})

	parent := make([]int, n)
	for i := range n {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	countComponents := func() int {
		seen := make(map[int]bool)
		for i := range n {
			r := find(i)
			seen[r] = true
		}
		return len(seen)
	}

	for _, p := range pairs {
		ra := find(p.a)
		rb := find(p.b)
		if ra != rb {
			parent[rb] = ra
			if countComponents() == 1 {
				return points[p.a].x * points[p.b].x
			}
		}
	}

	panic("impossible to have a single circuit")
}

func parseInput(input string) []string {
	s := strings.ReplaceAll(input, "\r", "")
	s = strings.TrimSpace(s)
	return strings.Split(s, "\n")
}

func parsePoints(input string) []point {
	lines := parseInput(input)
	points := make([]point, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(strings.TrimSpace(line), ",")
		x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		z, _ := strconv.Atoi(strings.TrimSpace(parts[2]))
		points = append(points, point{x: x, y: y, z: z})
	}
	return points
}

func squaredDistance(a, b point) int64 {
	dx := int64(a.x - b.x)
	dy := int64(a.y - b.y)
	dz := int64(a.z - b.z)
	return dx*dx + dy*dy + dz*dz
}
