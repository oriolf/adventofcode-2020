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
	util.Solve(solve1, solve2)
}

func solve1(lines []string) interface{} {
	rules, inputs := parseInput(lines)
	var count int
	for _, x := range inputs {
		if s, ok := matches(x, 0, rules); ok && s == "" {
			count++
		}
	}

	return count
}

func solve2(lines []string) interface{} {
	rules, inputs := parseInput(lines)
	rules[8] = rule{subrules: [][]int{[]int{42}, []int{42, 8}}}
	rules[11] = rule{subrules: [][]int{[]int{42, 31}, []int{42, 11, 31}}}

	var count int
	for _, x := range inputs {
		if ss := matches2([]string{x}, 0, rules); containsEmptyString(ss) {
			count++
		}
	}

	return count
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
		r.subrules = append(r.subrules, parseSubrules(strings.TrimSpace(x)))
	}

	return index, r
}

func parseSubrules(s string) (out []int) {
	for _, x := range strings.Split(s, " ") {
		out = append(out, util.ParseInt(x))
	}
	return out
}

func matches(s string, index int, rules map[int]rule) (string, bool) {
	r := rules[index]
	if r.literal != "" {
		if strings.HasPrefix(s, r.literal) {
			return s[len(r.literal):], true
		}
		return s, false
	}

	for _, l := range r.subrules {
		x, ok := matchesSubrules(s, l, rules)
		if ok {
			return x, true
		}
	}

	return s, false
}

func matchesSubrules(s string, l []int, rules map[int]rule) (string, bool) {
	x := s
	ok := true

	for _, i := range l {
		x, ok = matches(x, i, rules)
		if !ok {
			return s, false
		}
	}

	return x, true
}

func matches2(s []string, index int, rules map[int]rule) []string {
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
		out = append(out, matchesSubrules2(s, l, rules)...)
	}

	return out
}

func matchesSubrules2(s []string, l []int, rules map[int]rule) []string {
	for _, i := range l {
		s = matches2(s, i, rules)
	}

	return s
}
