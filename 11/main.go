package main

import (
	"github.com/oriolf/adventofcode2020/util"
	"strings"
)

func main() {
	util.Solve(solve(countOccupiedAround1, 4), solve(countOccupiedAround2, 5))
}

func solve(countFunc func([][]string, int, int) int, c int) func([]string) interface{} {
	return func(lines []string) interface{} {
		var seats [][]string
		for _, l := range lines {
			seats = append(seats, strings.Split(l, ""))
		}

		changed := true
		for changed {
			seats, changed = applyChanges(seats, countFunc, c)
		}

		return occupied(seats)
	}
}

func applyChanges(seats [][]string, countFunc func([][]string, int, int) int, c int) ([][]string, bool) {
	var changed bool
	counts := countOccupiedSeats(seats, countFunc)
	for i, l := range seats {
		for j, x := range l {
			if x == "L" && counts[i][j] == 0 {
				seats[i][j] = "#"
				changed = true
			} else if x == "#" && counts[i][j] >= c {
				seats[i][j] = "L"
				changed = true
			}
		}
	}

	return seats, changed
}

func countOccupiedSeats(seats [][]string, countFunc func([][]string, int, int) int) (out [][]int) {
	for i, l := range seats {
		o := make([]int, 0, len(l))
		for j := range l {
			o = append(o, countFunc(seats, i, j))
		}
		out = append(out, o)
	}

	return out
}

func countOccupiedAround1(seats [][]string, x, y int) (count int) {
	for i := util.Max(x-1, 0); i <= util.Min(x+1, len(seats)-1); i++ {
		for j := util.Max(y-1, 0); j <= util.Min(y+1, len(seats[0])-1); j++ {
			if i != x || j != y {
				if seats[i][j] == "#" {
					count++
				}
			}
		}
	}

	return count
}

func countOccupiedAround2(seats [][]string, x, y int) (count int) {
	count += countOccupiedDirection(seats, x, y, -1, -1)
	count += countOccupiedDirection(seats, x, y, -1, 0)
	count += countOccupiedDirection(seats, x, y, -1, 1)
	count += countOccupiedDirection(seats, x, y, 1, 1)
	count += countOccupiedDirection(seats, x, y, 1, 0)
	count += countOccupiedDirection(seats, x, y, 1, -1)
	count += countOccupiedDirection(seats, x, y, 0, 1)
	count += countOccupiedDirection(seats, x, y, 0, -1)
	return count
}

func countOccupiedDirection(seats [][]string, x, y, vx, vy int) int {
	i, j := x+vx, y+vy
	for i < len(seats) && j < len(seats[0]) && i >= 0 && j >= 0 {
		if seats[i][j] == "L" {
			return 0
		} else if seats[i][j] == "#" {
			return 1
		}
		i += vx
		j += vy
	}
	return 0
}

func occupied(seats [][]string) (count int) {
	for _, l := range seats {
		for _, x := range l {
			if x == "#" {
				count++
			}
		}
	}

	return count
}
