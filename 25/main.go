package main

import (
	"github.com/oriolf/adventofcode2020/util"
)

func main() {
	util.Solve(solve1, nil)
}

func solve1(lines []string) interface{} {
	cardPublic := util.ParseInt(lines[0])
	doorPublic := util.ParseInt(lines[1])
	doorLoop := getLoopSize(doorPublic)
	return loop(cardPublic, doorLoop)
}

func getLoopSize(n int) int {
	s := 1
	for i := 1; true; i++ {
		s *= 7
		s = s % 20201227
		if s == n {
			return i
		}
	}
	return 0
}

func loop(key, n int) int {
	s := 1
	for i := 0; i < n; i++ {
		s *= key
		s = s % 20201227
	}
	return s
}
