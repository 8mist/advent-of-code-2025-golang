package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
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

type MachineData struct {
	TargetLights  []int
	TargetJoltage []int
	Buttons       [][]int
}

func part1(input string) int {
	lines := parseInput(input)
	machines := parseMachines(lines)
	total := 0

	for _, m := range machines {
		if len(m.TargetLights) == 0 {
			continue
		}
		res, err := solveLights10(m)
		if err != nil {
			panic(err)
		}
		total += res
	}
	return total
}

func part2(input string) int {
	lines := parseInput(input)
	machines := parseMachines(lines)
	total := 0

	for _, m := range machines {
		if len(m.TargetJoltage) == 0 {
			continue
		}
		res, err := solveJoltage10(m)
		if err != nil {
			panic(err)
		}
		total += res
	}
	return total
}

func parseInput(input string) []string {
	s := strings.ReplaceAll(input, "\r", "")
	s = strings.TrimSpace(s)
	return strings.Split(s, "\n")
}

func parseList(s string) []int {
	s = strings.TrimSpace(s)
	if len(s) < 2 {
		return nil
	}
	s = s[1 : len(s)-1] // remove brackets

	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	result := make([]int, 0, len(parts))
	for _, p := range parts {
		val, err := strconv.Atoi(strings.TrimSpace(p))
		if err == nil {
			result = append(result, val)
		}
	}
	return result
}

func parseMachines(lines []string) []MachineData {
	machines := []MachineData{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		startBracket := strings.Index(line, "[")
		endBracket := strings.Index(line, "]")
		if startBracket == -1 || endBracket == -1 {
			continue
		}

		lightStr := line[startBracket+1 : endBracket]
		lights := make([]int, len(lightStr))
		for i, char := range lightStr {
			if char == '#' {
				lights[i] = 1
			} else {
				lights[i] = 0
			}
		}

		startBrace := strings.Index(line, "{")
		endBrace := strings.Index(line, "}")
		var joltage []int
		if startBrace != -1 && endBrace != -1 {
			joltage = parseList(line[startBrace : endBrace+1])
		}

		midSection := line[endBracket+1:]
		if startBrace != -1 {
			midSection = line[endBracket+1 : startBrace]
		}

		buttons := [][]int{}
		for {
			pStart := strings.Index(midSection, "(")
			if pStart == -1 {
				break
			}
			pEnd := strings.Index(midSection, ")")
			if pEnd == -1 {
				break
			}
			buttons = append(buttons, parseList(midSection[pStart:pEnd+1]))
			midSection = midSection[pEnd+1:]
		}

		machines = append(machines, MachineData{
			TargetLights:  lights,
			TargetJoltage: joltage,
			Buttons:       buttons,
		})
	}

	return machines
}

func solveLights10(m MachineData) (int, error) {
	N := len(m.TargetLights)
	M := len(m.Buttons)
	if N == 0 || M == 0 {
		return 0, nil
	}

	mat := make([][]int, N)
	for i := 0; i < N; i++ {
		row := make([]int, M+1)
		row[M] = m.TargetLights[i]
		mat[i] = row
	}

	for j, btn := range m.Buttons {
		for _, idx := range btn {
			if idx < N {
				mat[idx][j] = 1
			}
		}
	}

	pivotRow := 0
	pivotCols := map[int]int{}

	for col := 0; col < M && pivotRow < N; col++ {
		sel := -1
		for r := pivotRow; r < N; r++ {
			if mat[r][col] == 1 {
				sel = r
				break
			}
		}
		if sel == -1 {
			continue
		}

		mat[pivotRow], mat[sel] = mat[sel], mat[pivotRow]
		pivotCols[col] = pivotRow

		for r := 0; r < N; r++ {
			if r != pivotRow && mat[r][col] == 1 {
				for k := col; k <= M; k++ {
					mat[r][k] ^= mat[pivotRow][k]
				}
			}
		}
		pivotRow++
	}

	freeVars := []int{}
	for c := 0; c < M; c++ {
		if _, ok := pivotCols[c]; !ok {
			freeVars = append(freeVars, c)
		}
	}

	minPresses := M + 1
	count := 1 << len(freeVars)

	for mask := 0; mask < count; mask++ {
		x := make([]int, M)

		for i, c := range freeVars {
			if (mask>>i)&1 == 1 {
				x[c] = 1
			}
		}

		for c := M - 1; c >= 0; c-- {
			r, isPivot := pivotCols[c]
			if !isPivot {
				continue
			}

			val := mat[r][M]
			for k := c + 1; k < M; k++ {
				if mat[r][k] == 1 {
					val ^= x[k]
				}
			}
			x[c] = val
		}

		presses := 0
		for _, v := range x {
			presses += v
		}
		if presses < minPresses {
			minPresses = presses
		}
	}

	return minPresses, nil
}

func solveJoltage10(m MachineData) (int, error) {
	N := len(m.TargetJoltage)
	M := len(m.Buttons)
	if N == 0 || M == 0 {
		return 0, nil
	}

	mat := make([][]float64, N)
	for i := 0; i < N; i++ {
		row := make([]float64, M+1)
		row[M] = float64(m.TargetJoltage[i])
		mat[i] = row
	}

	for j, btn := range m.Buttons {
		for _, idx := range btn {
			if idx < N {
				mat[idx][j] = 1
			}
		}
	}

	pivotRow := 0
	pivotCols := map[int]int{}

	for col := 0; col < M && pivotRow < N; col++ {
		sel := -1
		for r := pivotRow; r < N; r++ {
			if math.Abs(mat[r][col]) > 1e-9 {
				sel = r
				break
			}
		}
		if sel == -1 {
			continue
		}

		mat[pivotRow], mat[sel] = mat[sel], mat[pivotRow]
		pivotCols[col] = pivotRow

		div := mat[pivotRow][col]
		for k := col; k <= M; k++ {
			mat[pivotRow][k] /= div
		}

		for r := 0; r < N; r++ {
			if r == pivotRow {
				continue
			}
			f := mat[r][col]
			if math.Abs(f) < 1e-9 {
				continue
			}
			for k := col; k <= M; k++ {
				mat[r][k] -= f * mat[pivotRow][k]
			}
		}

		pivotRow++
	}

	for r := pivotRow; r < N; r++ {
		if math.Abs(mat[r][M]) > 1e-9 {
			return 0, fmt.Errorf("inconsistent")
		}
	}

	freeVars := []int{}
	isPivotCol := make([]bool, M)
	for c := range pivotCols {
		isPivotCol[c] = true
	}
	for c := 0; c < M; c++ {
		if !isPivotCol[c] {
			freeVars = append(freeVars, c)
		}
	}

	maxTarget := 0.0
	for _, v := range m.TargetJoltage {
		if float64(v) > maxTarget {
			maxTarget = float64(v)
		}
	}
	searchBound := int(maxTarget) + 1

	minTotal := math.MaxInt64
	x := make([]float64, M)

	var dfs func(int, int)
	dfs = func(idx int, currentSum int) {
		if currentSum >= minTotal {
			return
		}
		if idx == len(freeVars) {
			total := currentSum
			valid := true

			for c := 0; c < M; c++ {
				r, isPivot := pivotCols[c]
				if !isPivot {
					continue
				}

				val := mat[r][M]
				for k := c + 1; k < M; k++ {
					if math.Abs(mat[r][k]) > 1e-9 {
						val -= mat[r][k] * x[k]
					}
				}

				if val < -1e-6 {
					valid = false
					break
				}

				rnd := math.Round(val)
				if math.Abs(val-rnd) > 1e-6 {
					valid = false
					break
				}

				iVal := int(rnd)
				x[c] = float64(iVal)
				total += iVal
			}

			if valid && total < minTotal {
				minTotal = total
			}

			return
		}

		c := freeVars[idx]
		for v := 0; v <= searchBound; v++ {
			x[c] = float64(v)
			dfs(idx+1, currentSum+v)
			if currentSum+v >= minTotal {
				break
			}
		}
	}

	dfs(0, 0)

	if minTotal == math.MaxInt64 {
		return 0, fmt.Errorf("no integer solution")
	}
	return minTotal, nil
}
