package main

import (
	"github.com/oriolf/adventofcode2020/util"
)

func main() {
	util.Solve(solve(sum1), solve(sum2))
}

func solve(sumFunc func([]map[string]int, []int) int) func([]string) interface{} {
	return func(lines []string) interface{} {
		group := map[string]int{}
		groupTotal := 0
		var groups []map[string]int
		var groupTotals []int
		for _, l := range lines {
			if l == "" {
				groups = append(groups, group)
				groupTotals = append(groupTotals, groupTotal)
				group = map[string]int{}
				groupTotal = 0
			} else {
				groupTotal++
				group = updateGroup(group, l)
			}
		}

		groups = append(groups, group)
		groupTotals = append(groupTotals, groupTotal)
		return sumFunc(groups, groupTotals)
	}
}

func sum1(groups []map[string]int, groupTotals []int) int {
	var sum int
	for _, group := range groups {
		sum += len(group)
	}
	return sum
}

func sum2(groups []map[string]int, groupTotals []int) int {
	var sum int
	for i, group := range groups {
		for _, v := range group {
			if v == groupTotals[i] {
				sum++
			}
		}
	}

	return sum
}

func updateGroup(group map[string]int, text string) map[string]int {
	for i := range text {
		group[text[i:i+1]]++
	}
	return group
}
