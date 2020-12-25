package main

import (
	"github.com/oriolf/adventofcode2020/util"
)

func main() {
	util.Solve(solve1, solve2)
}

func solve1(lines []string) interface{} {
	numbers := util.ParseInts(lines)
	x, y := sumPair(numbers, 2020)
	return x * y
}

func solve2(lines []string) interface{} {
	numbers := util.ParseInts(lines)
	for i, n := range numbers[1:] {
		x, y := sumPair(numbers[:i], 2020-n)
		if x+y == 2020-n {
			return x * y * n
		}
	}
	return 0
}

func sumPair(numbers []int, target int) (int, int) {
	seen := map[int]struct{}{}
	for _, n := range numbers {
		if _, ok := seen[target-n]; ok {
			return n, target - n
		} else {
			seen[n] = struct{}{}
		}
	}
	return 0, 0
}
