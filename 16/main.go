package main

import (
	"github.com/oriolf/adventofcode2020/util"
	"strings"
)

type rule struct {
	name           string
	range1, range2 [2]int
}

func (r rule) Matches(x int) bool {
	return (r.range1[0] <= x && x <= r.range1[1]) || (r.range2[0] <= x && x <= r.range2[1])
}

func main() {
	util.Solve(solve1, solve2)
}

func solve1(lines []string) interface{} {
	rules, _, otherTickets := parseTickets(lines)
	var sum int
	for _, t := range otherTickets {
		sum += countNotMatches(t, rules)
	}

	return sum
}

func solve2(lines []string) interface{} {
	rules, myTicket, otherTickets := parseTickets(lines)
	otherTickets = filterValid(otherTickets, rules)
	fieldsMap := getFieldsRuleName(otherTickets, rules)

	mult := 1
	for k, v := range fieldsMap {
		if strings.HasPrefix(v, "departure") {
			mult *= myTicket[k]
		}
	}

	return mult
}

func getFieldsRuleName(tickets [][]int, rules []rule) map[int]string {
	m := make(map[int][]rule)
	for i := range tickets[0] {
		m[i] = rules
	}

	for i := 0; i < len(tickets[0]); i++ {
		m[i] = removeNotMatching(rules, tickets, i)
	}

	var toRemove []rule
	for multipleRules(m) {
		m, toRemove = removeIncompatibles(m, toRemove)
	}

	out := make(map[int]string)
	for k, v := range m {
		out[k] = v[0].name
	}

	return out
}

func multipleRules(m map[int][]rule) bool {
	for _, x := range m {
		if len(x) > 1 {
			return true
		}
	}
	return false
}

func removeIncompatibles(m map[int][]rule, toRemove []rule) (map[int][]rule, []rule) {
	for _, v := range m {
		if len(v) == 1 && !ruleInList(v[0], toRemove) {
			toRemove = append(toRemove, v[0])
		}
	}

	for k, v := range m {
		m[k] = removeRules(v, toRemove)
	}

	return m, toRemove
}

func removeRules(in, toRemove []rule) (out []rule) {
	if len(in) == 1 {
		return in
	}
	for _, r := range in {
		if !ruleInList(r, toRemove) {
			out = append(out, r)
		}
	}
	return out
}

func ruleInList(r rule, l []rule) bool {
	for _, x := range l {
		if x == r {
			return true
		}
	}
	return false
}

func removeNotMatching(rules []rule, tickets [][]int, i int) (out []rule) {
LOOP:
	for _, r := range rules {
		for _, t := range tickets {
			if !r.Matches(t[i]) {
				continue LOOP
			}
		}
		out = append(out, r)
	}
	return out
}

func filterValid(tickets [][]int, rules []rule) (out [][]int) {
LOOP:
	for _, t := range tickets {
		for _, x := range t {
			if matchesNoRule(x, rules) {
				continue LOOP
			}
		}
		out = append(out, t)
	}

	return out
}

func matchesNoRule(x int, rules []rule) bool {
	for _, r := range rules {
		if r.Matches(x) {
			return false
		}
	}
	return true
}

func countNotMatches(t []int, rules []rule) (sum int) {
LOOP:
	for _, x := range t {
		for _, r := range rules {
			if r.Matches(x) {
				continue LOOP
			}
		}

		sum += x
	}

	return sum
}

func parseTickets(lines []string) (rules []rule, t []int, tickets [][]int) {
	var index int
	for _, l := range lines {
		index++
		if l == "" {
			break
		}
		rules = append(rules, parseRule(l))
	}

	index++
	t = parseTicket(lines[index])
	index += 3
	for _, l := range lines[index:] {
		tickets = append(tickets, parseTicket(l))
	}

	return rules, t, tickets
}

func parseTicket(s string) (out []int) {
	for _, x := range strings.Split(s, ",") {
		out = append(out, util.ParseInt(x))
	}
	return out
}

func parseRule(s string) rule {
	parts := strings.Split(s, ":")
	ranges := strings.Split(parts[1], " or ")

	return rule{name: parts[0], range1: parseRange(ranges[0]), range2: parseRange(ranges[1])}
}

func parseRange(s string) [2]int {
	parts := strings.Split(s, "-")
	return [2]int{util.ParseInt(parts[0]), util.ParseInt(parts[1])}
}
