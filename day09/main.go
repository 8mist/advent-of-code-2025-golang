package main

import (
	"8mist/aoc/mathy"
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

type Point struct {
	x, y int
}

type Edge struct {
	x1, y1 int
	x2, y2 int
	hor    bool
}

func part1(input string) int {
	lines := parseInput(input)
	points := parsePoints(lines)

	maxArea := 0
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			maxArea = max(maxArea, (mathy.Abs(points[i].x-points[j].x)+1)*(mathy.Abs(points[i].y-points[j].y)+1))
		}
	}
	return maxArea
}

func part2(input string) int {
	lines := parseInput(input)
	points := parsePoints(lines)
	n := len(points)
	edges := buildEdges(points)

	best := 0
	for i := range n {
		a := points[i]
		for j := i + 1; j < n; j++ {
			b := points[j]

			x1 := min(a.x, b.x)
			x2 := max(a.x, b.x)
			y1 := min(a.y, b.y)
			y2 := max(a.y, b.y)

			dx := x2 - x1 + 1
			dy := y2 - y1 + 1
			area := dx * dy

			if area <= best {
				continue
			}

			c3 := Point{x1, y2}
			c4 := Point{x2, y1}

			if !pointInsideOrOn(c3, points, edges) || !pointInsideOrOn(c4, points, edges) {
				continue
			}
			if rectangleCutByPolygon(x1, y1, x2, y2, edges) {
				continue
			}

			if area > best {
				best = area
			}
		}
	}

	return best
}

func parseInput(input string) []string {
	s := strings.ReplaceAll(input, "\r", "")
	s = strings.TrimSpace(s)
	return strings.Split(s, "\n")
}

func parsePoints(lines []string) []Point {
	points := make([]Point, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		points = append(points, Point{x: x, y: y})
	}
	return points
}

func buildEdges(points []Point) []Edge {
	n := len(points)
	edges := make([]Edge, 0, n)
	for i := range n {
		a := points[i]
		b := points[(i+1)%n]
		e := Edge{x1: a.x, y1: a.y, x2: b.x, y2: b.y}
		if a.y == b.y {
			e.hor = true
			if e.x1 > e.x2 {
				e.x1, e.x2 = e.x2, e.x1
			}
		} else {
			e.hor = false
			if e.y1 > e.y2 {
				e.y1, e.y2 = e.y2, e.y1
			}
		}
		edges = append(edges, e)
	}
	return edges
}

func pointInsideOrOn(p Point, points []Point, edges []Edge) bool {
	for _, e := range edges {
		if e.hor {
			if p.y == e.y1 && p.x >= e.x1 && p.x <= e.x2 {
				return true
			}
		} else {
			if p.x == e.x1 && p.y >= e.y1 && p.y <= e.y2 {
				return true
			}
		}
	}
	return pointInPolygonRayCast(p, points)
}

func pointInPolygonRayCast(p Point, poly []Point) bool {
	inside := false
	n := len(poly)
	if n < 3 {
		return false
	}

	j := n - 1
	for i := range n {
		pi := poly[i]
		pj := poly[j]

		if (pi.y > p.y) != (pj.y > p.y) {
			xIntersect := float64(pj.x) + float64(p.y-pj.y)*float64(pi.x-pj.x)/float64(pi.y-pj.y)
			if float64(p.x) < xIntersect {
				inside = !inside
			}
		}
		j = i
	}
	return inside
}

func rectangleCutByPolygon(x1, y1, x2, y2 int, edges []Edge) bool {
	if x1 == x2 || y1 == y2 {
		return false
	}

	for _, e := range edges {
		if e.hor {
			y0 := e.y1
			if y0 <= y1 || y0 >= y2 {
				continue
			}
			if max(e.x1, x1) < min(e.x2, x2) {
				return true
			}
		} else {
			x0 := e.x1
			if x0 <= x1 || x0 >= x2 {
				continue
			}
			if max(e.y1, y1) < min(e.y2, y2) {
				return true
			}
		}
	}
	return false
}
