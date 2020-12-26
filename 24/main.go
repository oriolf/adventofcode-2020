package main

import (
	"fmt"
	"github.com/oriolf/adventofcode2020/util"
	"strings"
)

func main() {
	util.Solve(solve1, solve2)
}

func solve1(lines []string) interface{} {
	grid := initializeGrid(lines)
	return countBlacks(grid)
}

func solve2(lines []string) interface{} {
	grid := initializeGrid(lines)
	for i := 0; i < 100; i++ {
		grid = step(growGrid(grid))
	}

	return countBlacks(grid)
}

func initializeGrid(lines []string) [][]bool {
	movements := parseMovements(lines)
	grid, cx, cy := generateGrid(movements)
	for _, movement := range movements {
		grid = flipTile(grid, cx, cy, movement)
	}

	return grid
}

func step(grid [][]bool) [][]bool {
	counts := countAround(grid)
	for i, row := range grid {
		for j := range row {
			if grid[i][j] { // black
				if counts[i][j] == 0 || counts[i][j] > 2 {
					grid[i][j] = false
				}
			} else { // white
				if counts[i][j] == 2 {
					grid[i][j] = true
				}
			}
		}
	}

	return grid
}

func growGrid(in [][]bool) (out [][]bool) {
	out = append(out, util.GenerateBoolRow(len(in[0])+2))
	for _, row := range in {
		out = append(out, growRow(row))
	}
	out = append(out, util.GenerateBoolRow(len(in[0])+2))
	return out
}

func growRow(in []bool) (out []bool) {
	out = append([]bool{false}, in...)
	return append(out, false)
}

func countAround(grid [][]bool) [][]int {
	counts := util.GenerateIntMatrix(len(grid), len(grid[0]))
	for i, row := range grid {
		for j := range row {
			counts[i][j] = countAroundSingle(grid, i, j)
		}
	}

	return counts
}

func countAroundSingle(grid [][]bool, i, j int) (count int) {
	for _, v := range [][2]int{
		[2]int{1, 1},
		[2]int{0, -1},
		[2]int{0, 1},
		[2]int{-1, -1},
		[2]int{1, 0},
		[2]int{-1, 0},
	} {
		x, y := i+v[0], j+v[1]
		if x > 0 && y > 0 && x < len(grid) && y < len(grid[0]) {
			if grid[x][y] {
				count++
			}
		}
	}

	return count
}

func parseMovements(lines []string) (movements [][][2]int) {
	for _, l := range lines {
		movements = append(movements, parseMovementsLine(l))
	}
	return movements
}

func parseMovementsLine(l string) (movements [][2]int) {
	for len(l) > 0 {
		if strings.HasPrefix(l, "ne") {
			movements = append(movements, [2]int{1, 1})
			l = l[2:]
		} else if strings.HasPrefix(l, "se") {
			movements = append(movements, [2]int{0, -1})
			l = l[2:]
		} else if strings.HasPrefix(l, "nw") {
			movements = append(movements, [2]int{0, 1})
			l = l[2:]
		} else if strings.HasPrefix(l, "sw") {
			movements = append(movements, [2]int{-1, -1})
			l = l[2:]
		} else if strings.HasPrefix(l, "e") {
			movements = append(movements, [2]int{1, 0})
			l = l[1:]
		} else if strings.HasPrefix(l, "w") {
			movements = append(movements, [2]int{-1, 0})
			l = l[1:]
		} else {
			panic(fmt.Sprintf("unknown input %q", l))
		}
	}

	return movements
}

func countBlacks(in [][]bool) (count int) {
	for _, x := range in {
		for _, y := range x {
			if y {
				count++
			}
		}
	}

	return count
}

func flipTile(grid [][]bool, cx, cy int, movements [][2]int) [][]bool {
	x, y := finalPoint(cx, cy, movements)
	grid[x][y] = !grid[x][y]
	return grid
}

func finalPoint(cx, cy int, movements [][2]int) (int, int) {
	for _, movement := range movements {
		cx, cy = cx+movement[0], cy+movement[1]
	}
	return cx, cy
}

func generateGrid(allMovements [][][2]int) ([][]bool, int, int) {
	var maxx, minx, maxy, miny int
	for _, movements := range allMovements {
		x, y := finalPoint(0, 0, movements)
		minx = util.Min(minx, x)
		maxx = util.Max(maxx, x)
		miny = util.Min(miny, y)
		maxy = util.Max(maxy, y)
	}

	return util.GenerateBoolMatrix(maxx-minx+1, maxy-miny+1), -minx, maxy
}
