package main

import (
	"github.com/oriolf/adventofcode2020/util"
	"strings"
)

func main() {
	util.Solve(solve1, nil)
}

func solve1(lines []string) interface{} {
	turnsSpoken := make(map[int][]int)
	turn := 1
	var lastSpoken int
	for _, x := range strings.Split(lines[0], ",") {
		y := util.ParseInt(x)
		turnsSpoken[y] = append(turnsSpoken[y], turn)
		lastSpoken = y
		turn++
	}

	for turn <= 30000000 { // 2020 {
		lastSpoken = computeNextSpoken(turnsSpoken, lastSpoken)
		turnsSpoken[lastSpoken] = append(turnsSpoken[lastSpoken], turn)
		turn++
	}

	return lastSpoken
}

func computeNextSpoken(turnsSpoken map[int][]int, lastSpoken int) int {
	if len(turnsSpoken[lastSpoken]) == 1 {
		return 0
	}

	l := turnsSpoken[lastSpoken]
	return l[len(l)-1] - l[len(l)-2]
}
