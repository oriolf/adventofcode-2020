package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bag struct {
	name       string
	canContain map[string]int
}

func main() {
	bags := map[string]bag{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		bag := parseBag(scanner.Text())
		bags[bag.name] = bag
	}

	target := "shiny gold"
	//	var count int
	//	for _, b := range bags {
	//		if canContain(b, bags, target) {
	//			count++
	//		}
	//	}
	//
	//	fmt.Println(count)

	fmt.Println(countBags(bags, target) - 1)
}

func countBags(bags map[string]bag, target string) int {
	b, ok := bags[target]
	if !ok {
		panic("no bag")
	}
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
			n, err := strconv.Atoi(fields[i-1])
			if err != nil {
				panic("bad number")
			}
			l[strings.Join(fields[i:i+2], " ")] = n
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
