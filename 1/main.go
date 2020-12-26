package main

import (
	"github.com/oriolf/adventofcode2020/util"
)

func main() {
	util.Solve(solve1, solve2)
}

func solve1(lines []string) interface{} {
	numbers := util.ParseInts(lines)
	x, y := util.SumPair(numbers, 2020)
	return x * y
}

func solve2(lines []string) interface{} {
	numbers := util.ParseInts(lines)
	for i, n := range numbers[1:] {
		x, y := util.SumPair(numbers[:i], 2020-n)
		if x+y == 2020-n {
			return x * y * n
		}
	}
	return 0
}
