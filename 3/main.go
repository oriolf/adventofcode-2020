package main

import (
	"github.com/oriolf/adventofcode2020/util"
)

func main() {
	util.Solve(solve1, solve2)
}

func solve1(lines []string) interface{} {
	forest := util.ParseBoolMatrix(lines)
	return countDescent(forest, 3, 1)
}

func solve2(lines []string) interface{} {
	forest := util.ParseBoolMatrix(lines)
	result := 1
	for _, x := range [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	} {
		result *= countDescent(forest, x[0], x[1])
	}

	return result
}

func countDescent(forest [][]bool, right, down int) (count int) {
	coln := len(forest[0])
	row, column := down, right
	for row < len(forest) {
		if forest[row][column%coln] {
			count++
		}
		row += down
		column += right
	}
	return count
}
