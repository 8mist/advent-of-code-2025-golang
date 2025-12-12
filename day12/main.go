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
	}
}

type point struct {
	x, y int
}

type variant struct {
	width, height int
	cells         []point
}

type shape struct {
	area     int
	variants []variant
}

type region struct {
	width, height int
	counts        []int
}

func part1(input string) int {
	shapes, regions := parseInput(input)
	valid := 0
	for _, r := range regions {
		if regionCanFit(shapes, r) {
			valid++
		}
	}
	return valid
}

func parseInput(input string) ([]shape, []region) {
	lines := strings.Split(strings.ReplaceAll(input, "\r", ""), "\n")
	shapes := []shape{}
	regions := []region{}

	// --- Parse shape blocks of form:
	// 0:
	// ###
	// ..#
	// ###
	i := 0
	for i < len(lines) {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			i++
			continue
		}

		if isRegionLine(line) {
			break
		}

		// Expect a header like "0:" or "5:"
		if !strings.HasSuffix(line, ":") {
			i++
			continue
		}
		i++

		var rows []string
		for i < len(lines) {
			s := strings.TrimSpace(lines[i])
			if s == "" || strings.HasSuffix(s, ":") || isRegionLine(s) {
				break
			}
			rows = append(rows, s)
			i++
		}
		if len(rows) > 0 {
			shapes = append(shapes, buildShape(rows))
		}
	}

	// --- Parse regions: "WxH: c0 c1 c2 ..."
	for i < len(lines) {
		line := strings.TrimSpace(lines[i])
		i++
		if !isRegionLine(line) {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		dimPart := strings.TrimSpace(parts[0])
		cntPart := strings.TrimSpace(parts[1])

		wh := strings.Split(dimPart, "x")
		if len(wh) != 2 {
			continue
		}
		w, _ := strconv.Atoi(strings.TrimSpace(wh[0]))
		h, _ := strconv.Atoi(strings.TrimSpace(wh[1]))

		countFields := strings.Fields(cntPart)
		counts := make([]int, len(countFields))
		for idx, s := range countFields {
			counts[idx], _ = strconv.Atoi(s)
		}

		regions = append(regions, region{
			width:  w,
			height: h,
			counts: counts,
		})
	}
	return shapes, regions
}

func isRegionLine(s string) bool {
	s = strings.TrimSpace(s)
	colonIdx := strings.IndexByte(s, ':')
	if colonIdx <= 0 {
		return false
	}
	head := s[:colonIdx]
	parts := strings.Split(head, "x")
	if len(parts) != 2 {
		return false
	}
	if _, err := strconv.Atoi(strings.TrimSpace(parts[0])); err != nil {
		return false
	}
	if _, err := strconv.Atoi(strings.TrimSpace(parts[1])); err != nil {
		return false
	}
	return true
}

func buildShape(rows []string) shape {
	h0 := len(rows)
	w0 := 0
	for _, row := range rows {
		if len(row) > w0 {
			w0 = len(row)
		}
	}
	grid := make([][]bool, h0)
	for y, row := range rows {
		grid[y] = make([]bool, w0)
		for x := 0; x < len(row); x++ {
			if row[x] == '#' {
				grid[y][x] = true
			}
		}
	}

	var variants []variant
	seen := make(map[string]bool)

	g := grid
	for r := 0; r < 4; r++ {
		if r > 0 {
			g = rotateGrid(g)
		}
		for f := 0; f < 2; f++ {
			gf := g
			if f == 1 {
				gf = flipGridH(g)
			}
			v := gridToVariant(gf)
			if len(v.cells) == 0 {
				continue
			}
			key := variantKey(v)
			if !seen[key] {
				seen[key] = true
				variants = append(variants, v)
			}
		}
	}

	area := 0
	if len(variants) > 0 {
		area = len(variants[0].cells)
	}

	return shape{area: area, variants: variants}
}

func rotateGrid(grid [][]bool) [][]bool {
	height := len(grid)
	if height == 0 {
		return grid
	}
	width := len(grid[0])
	res := make([][]bool, width)
	for i := range res {
		res[i] = make([]bool, height)
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			res[width-1-x][y] = grid[y][x]
		}
	}
	return res
}

func flipGridH(grid [][]bool) [][]bool {
	height := len(grid)
	if height == 0 {
		return grid
	}
	width := len(grid[0])
	res := make([][]bool, height)
	for y := 0; y < height; y++ {
		res[y] = make([]bool, width)
		for x := 0; x < width; x++ {
			res[y][width-1-x] = grid[y][x]
		}
	}
	return res
}

func gridToVariant(grid [][]bool) variant {
	height := len(grid)
	if height == 0 {
		return variant{}
	}
	width := len(grid[0])

	minX, minY := width, height
	maxX, maxY := -1, -1

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] {
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}
	if maxX < minX || maxY < minY {
		return variant{}
	}

	vw := maxX - minX + 1
	vh := maxY - minY + 1
	var cells []point
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if grid[y][x] {
				cells = append(cells, point{x: x - minX, y: y - minY})
			}
		}
	}

	return variant{width: vw, height: vh, cells: cells}
}

func variantKey(v variant) string {
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(v.width))
	sb.WriteByte('x')
	sb.WriteString(strconv.Itoa(v.height))
	sb.WriteByte(':')
	for _, c := range v.cells {
		sb.WriteString(strconv.Itoa(c.x))
		sb.WriteByte(',')
		sb.WriteString(strconv.Itoa(c.y))
		sb.WriteByte(';')
	}
	return sb.String()
}

const smallBoardMaxArea12 = 15 * 15 // full tiling search only if w*h <= this

func regionCanFit(shapes []shape, r region) bool {
	if len(shapes) == 0 {
		return false
	}

	// Area check
	totalArea := 0
	for i, cnt := range r.counts {
		if i >= len(shapes) {
			break
		}
		totalArea += cnt * shapes[i].area
	}
	if totalArea > r.width*r.height {
		return false
	}

	// Small boards: do an actual tiling search (geometric fit).
	if r.width*r.height <= smallBoardMaxArea12 {
		return canTileRegionSmall(shapes, r)
	}
	return true
}

func canTileRegionSmall(shapes []shape, r region) bool {
	w, h := r.width, r.height
	numShapes := len(shapes)

	// Precompute all possible placements for each shape type on this board.
	allPlacements := make([][][]int, numShapes)
	for si := 0; si < numShapes; si++ {
		sh := shapes[si]
		for _, v := range sh.variants {
			if v.width == 0 || v.height == 0 {
				continue
			}
			for by := 0; by <= h-v.height; by++ {
				for bx := 0; bx <= w-v.width; bx++ {
					var cells []int
					for _, c := range v.cells {
						x, y := bx+c.x, by+c.y
						cells = append(cells, y*w+x)
					}
					allPlacements[si] = append(allPlacements[si], cells)
				}
			}
		}
	}

	board := make([]bool, w*h)
	counts := make([]int, len(r.counts))
	copy(counts, r.counts)

	return btTile(shapes, board, w, h, counts, allPlacements)
}

func btTile(shapes []shape, board []bool, w, h int, counts []int, placements [][][]int) bool {
	// Check if all counts are zero (all presents placed).
	done := true
	totalRemainingArea := 0
	for i, c := range counts {
		if i >= len(shapes) {
			break
		}
		if c > 0 {
			done = false
			totalRemainingArea += c * shapes[i].area
		}
	}
	if done {
		return true
	}

	// Quick prune: not enough free cells to fit remaining area.
	freeCells := 0
	for _, b := range board {
		if !b {
			freeCells++
		}
	}
	if totalRemainingArea > freeCells {
		return false
	}

	// Choose the shape type that is currently most constrained:
	// the one with the smallest number of feasible placements.
	bestShape, bestCount := -1, 1<<30
	for si, c := range counts {
		if c <= 0 {
			continue
		}
		cnt := 0
		for _, pl := range placements[si] {
			ok := true
			for _, idx := range pl {
				if board[idx] {
					ok = false
					break
				}
			}
			if ok {
				cnt++
				if cnt >= bestCount {
					break
				}
			}
		}
		if cnt == 0 {
			return false
		}
		if cnt < bestCount {
			bestCount = cnt
			bestShape = si
		}
	}

	// Try placing one copy of bestShape in each feasible way.
	counts[bestShape]--
	for _, pl := range placements[bestShape] {
		ok := true
		for _, idx := range pl {
			if board[idx] {
				ok = false
				break
			}
		}
		if !ok {
			continue
		}

		// Place shape
		for _, idx := range pl {
			board[idx] = true
		}

		if btTile(shapes, board, w, h, counts, placements) {
			return true
		}

		// Undo
		for _, idx := range pl {
			board[idx] = false
		}
	}
	counts[bestShape]++
	return false
}
