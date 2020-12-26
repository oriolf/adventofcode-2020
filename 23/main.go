package main

import (
	"fmt"
	"github.com/oriolf/adventofcode2020/util"
	"strings"
)

type cup struct {
	next  *cup
	value int
}

func main() {
	util.Solve(solve1, solve2)
}

func solve1(lines []string) interface{} {
	circle := step(parseInput(lines), 100, 9)
	return formatLinked(circle)
}

func solve2(lines []string) interface{} {
	one := step(completeCircle(parseInput(lines)), 10000000, 1000000)
	return one.next.value * one.next.next.value
}

func parseInput(l []string) []int { return util.ParseInts(strings.Split(l[0], "")) }

func completeCircle(in []int) []int {
	for i := 10; i <= 1000000; i++ {
		in = append(in, i)
	}
	return in
}

func step(nums []int, n, m int) *cup {
	circle, valuesMap := parseLinked(nums)
	for i := 0; i < n; i++ {
		circle = stepLinked(circle, valuesMap, m)
	}
	return valuesMap[1]
}

func parseLinked(l []int) (*cup, []*cup) {
	m := make([]*cup, len(l)+1)
	originalHead := &cup{value: l[0]}
	m[l[0]] = originalHead
	head := originalHead
	for _, x := range l[1:] {
		c := &cup{value: x}
		m[x] = c
		head.next = c
		head = c
	}

	head.next = originalHead
	return originalHead, m
}

func stepLinked(circle *cup, valuesMap []*cup, max int) *cup {
	current := circle.value
	pickedUp, circle := pickUp(circle)
	destination := getDestination(current, max, pickedUp, circle)
	insert(circle, pickedUp, valuesMap[destination])
	return circle
}

func getDestination(current, max int, pickedUp, circle *cup) int {
LOOP:
	for {
		current--
		if current == 0 {
			current = max
		}
		if pickedUp.value == current ||
			pickedUp.next.value == current ||
			pickedUp.next.next.value == current {
			continue LOOP
		}
		return current
	}
}

func insert(circle, pickedUp, destination *cup) {
	pickedUp.next.next.next = destination.next
	destination.next = pickedUp
}

func formatLinked(circle *cup) (out string) {
	circle = circle.next
	for circle.value != 1 {
		out += fmt.Sprintf("%d", circle.value)
		circle = circle.next
	}
	return out
}

func pickUp(circle *cup) (picked *cup, out *cup) {
	picked = circle.next
	fourth := circle.next.next.next.next
	circle.next = fourth
	return picked, fourth
}
