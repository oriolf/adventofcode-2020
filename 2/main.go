package main

import (
	"fmt"
	"github.com/oriolf/adventofcode2020/util"
	"strings"
)

type testcase struct {
	min, max    int
	char, value string
}

func main() {
	util.Solve(solve(isOk1), solve(isOk2))
}

func solve(okFunc func(testcase) bool) func([]string) interface{} {
	return func(lines []string) interface{} {
		var count int
		for _, x := range lines {
			if okFunc(getCase(x)) {
				count++
			}
		}

		return count
	}
}

func getCase(s string) (out testcase) {
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, ":", "")
	if _, err := fmt.Sscanf(s, "%d %d %s %s", &out.min, &out.max, &out.char, &out.value); err != nil {
		panic(err)
	}
	return out
}

func isOk1(x testcase) bool {
	var count int
	for i := range x.value {
		if x.value[i:i+1] == x.char {
			count++
		}
	}
	return count >= x.min && count <= x.max
}

func isOk2(x testcase) bool {
	var count int
	for i := range x.value {
		if (i == x.min-1 || i == x.max-1) && x.value[i:i+1] == x.char {
			count++
		}
	}
	return count == 1
}
