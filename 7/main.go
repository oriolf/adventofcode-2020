package main

import (
	"github.com/oriolf/adventofcode2020/util"
	"strings"
)

type bag struct {
	name       string
	canContain map[string]int
}

var target = "shiny gold"

func main() {
	util.Solve(solve1, solve2)
}

func solve1(lines []string) interface{} {
	bags := parseBags(lines)
	var count int
	for _, b := range bags {
		if canContain(b, bags, target) {
			count++
		}
	}

	return count
}

func solve2(lines []string) interface{} {
	bags := parseBags(lines)
	return countBags(bags, target) - 1
}

func parseBags(lines []string) map[string]bag {
	bags := map[string]bag{}
	for _, l := range lines {
		bag := parseBag(l)
		bags[bag.name] = bag
	}
	return bags
}

func countBags(bags map[string]bag, target string) int {
	b := bags[target]
	count := 1
	for bb, m := range b.canContain {
		count += m * countBags(bags, bb)
	}
	return count
}

func parseBag(s string) bag {
	fields := strings.Fields(strings.TrimSpace(strings.Trim(s, ".")))
	name := strings.Join(fields[:2], " ")
	l := map[string]int{}
	fields = fields[4:]
	if len(fields) > 3 {
		for i := 1; i < len(fields); i += 4 {
			l[strings.Join(fields[i:i+2], " ")] = util.ParseInt(fields[i-1])
		}
	}

	return bag{name: name, canContain: l}
}

func canContain(b bag, bags map[string]bag, target string) bool {
	for cc, _ := range b.canContain {
		if cc == target {
			return true
		} else {
			bb, ok := bags[cc]
			if ok && canContain(bb, bags, target) {
				return true
			}
		}
	}

	return false
}
