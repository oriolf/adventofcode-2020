package main

import (
	"github.com/oriolf/adventofcode2020/util"
	"strings"
)

type rule struct {
	literal  string
	subrules [][]int
}

func main() {
	util.Solve(solve(rules1), solve(rules2))
}

func solve(rulesFunc func(map[int]rule) map[int]rule) func([]string) interface{} {
	return func(lines []string) interface{} {
		rules, inputs := parseInput(lines)
		rules = rulesFunc(rules)

		var count int
		for _, x := range inputs {
			if ss := matches([]string{x}, 0, rules); containsEmptyString(ss) {
				count++
			}
		}

		return count
	}
}

func rules1(rules map[int]rule) map[int]rule { return rules }
func rules2(rules map[int]rule) map[int]rule {
	rules[8] = rule{subrules: [][]int{[]int{42}, []int{42, 8}}}
	rules[11] = rule{subrules: [][]int{[]int{42, 31}, []int{42, 11, 31}}}
	return rules
}

func containsEmptyString(l []string) bool {
	for _, x := range l {
		if x == "" {
			return true
		}
	}
	return false
}

func parseInput(lines []string) (map[int]rule, []string) {
	rules := make(map[int]rule)
	var i int
	for _, x := range lines {
		i++
		if x == "" {
			break
		}
		index, r := parseRule(x)
		rules[index] = r
	}

	return rules, lines[i:]
}

func parseRule(s string) (int, rule) {
	parts := strings.Split(s, ": ")
	index := util.ParseInt(parts[0])
	s = parts[1]
	if strings.HasPrefix(s, `"`) {
		return index, rule{literal: strings.Trim(s, `"`)}
	}

	var r rule
	parts = strings.Split(s, "|")
	for _, x := range parts {
		l := strings.Split(strings.TrimSpace(x), " ")
		r.subrules = append(r.subrules, util.ParseInts(l))
	}

	return index, r
}

func matches(s []string, index int, rules map[int]rule) []string {
	r := rules[index]
	if len(s) == 0 {
		return nil
	}

	var out []string
	if r.literal != "" {
		for _, x := range s {
			if strings.HasPrefix(x, r.literal) {
				out = append(out, x[len(r.literal):])
			}
		}
		return out
	}

	for _, l := range r.subrules {
		out = append(out, matchesSubrules(s, l, rules)...)
	}

	return out
}

func matchesSubrules(s []string, l []int, rules map[int]rule) []string {
	for _, i := range l {
		s = matches(s, i, rules)
	}

	return s
}
