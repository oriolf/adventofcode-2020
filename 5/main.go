package main

import (
	"github.com/oriolf/adventofcode2020/util"
)

func main() {
	util.Solve(solve1, solve2)
}

func solve1(lines []string) interface{} {
	var max int
	present := parseBoardingPasses(lines)
	for k := range present {
		if k > max {
			max = k
		}
	}
	return max
}

func solve2(lines []string) interface{} {
	present := parseBoardingPasses(lines)
	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			res := i*8 + j
			if _, ok := present[res]; !ok {
				_, ok1 := present[res+1]
				_, ok2 := present[res-1]
				if ok1 && ok2 {
					return res
				}
			}
		}
	}

	return ""
}

func parseBoardingPasses(lines []string) map[int]struct{} {
	present := map[int]struct{}{}
	for _, l := range lines {
		row, column := parseBoardingPass(l)
		present[row*8+column] = struct{}{}
	}

	return present
}

func parseBoardingPass(s string) (int, int) {
	var row, column []bool
	for i := 0; i < 7; i++ {
		if s[i:i+1] == "B" {
			row = append(row, true)
		} else {
			row = append(row, false)
		}
	}

	for i := 7; i < 10; i++ {
		if s[i:i+1] == "R" {
			column = append(column, true)
		} else {
			column = append(column, false)
		}
	}

	return binToInt(row), binToInt(column)
}

func binToInt(bin []bool) (out int) {
	mult := 1
	for i := len(bin) - 1; i >= 0; i-- {
		if bin[i] {
			out += mult
		}
		mult *= 2
	}
	return out
}
