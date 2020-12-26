package main

import (
	"github.com/oriolf/adventofcode2020/util"
)

func main() {
	util.Solve(solve1, solve2)
}

func solve1(lines []string) interface{} {
	var preamble = 25
	if len(lines) < preamble {
		preamble = 5
	}
	const target = 127
	numbers := util.ParseInts(lines)
	for i := preamble; i < len(numbers); i++ {
		x, y := util.SumPair(numbers[i-preamble:i], numbers[i])
		if x == 0 && y == 0 {
			return numbers[i]
		}
	}
	return ""
}

func solve2(lines []string) interface{} {
	const target = 675280050
	numbers := util.ParseInts(lines)
	for i := 0; i < len(numbers); i++ {
		slice, ok := findSliceSum(numbers[i:], target)
		if ok {
			return util.ListMin(slice) + util.ListMax(slice)
		}
	}
	return ""
}

func findSliceSum(numbers []int, target int) (out []int, ok bool) {
	var sum int
	for _, x := range numbers {
		sum += x
		out = append(out, x)
		if sum == target && len(out) > 1 {
			return out, true
		} else if sum > target {
			return nil, false
		}
	}

	return nil, false
}
