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
	//	util.Solve(solve1, solve2) // naive
	util.Solve(solveQuick1, solveQuick2)
}

func solveQuick1(lines []string) interface{} {
	circle, valuesMap := parseLinked(parseInput(lines))
	for i := 0; i < 100; i++ {
		circle = stepLinked(circle, valuesMap, 9)
	}
	return formatLinked(circle)
}

func solveQuick2(lines []string) interface{} {
	circle, valuesMap := parseLinked(completeCircle(parseInput(lines)))
	for i := 0; i < 10000000; i++ {
		circle = stepLinked(circle, valuesMap, 1000000)
	}

	one := valuesMap[1]
	return one.next.value * one.next.next.value
}

func parseLinked(l []int) (*cup, map[int]*cup) {
	m := make(map[int]*cup)
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

func stepLinked(circle *cup, valuesMap map[int]*cup, max int) *cup {
	current := circle.value
	pickedUp, circle := pickUpLinked(circle)
	destination := getDestinationLinked(current, max, pickedUp, circle)
	insertLinked(circle, pickedUp, valuesMap[destination])
	return circle
}

func solve1(lines []string) interface{} {
	circle, index := parseInput(lines), 0
	for i := 0; i < 100; i++ {
		circle, index = step(circle, index)
	}
	return formatCircle(circle)
}

func solve2(lines []string) interface{} {
	circle, index := parseInput(lines), 0
	circle = completeCircle(circle)
	for i := 0; i < 10000000; i++ {
		if i%1000 == 0 {
			fmt.Println(i / 1000)
		}
		circle, index = step(circle, index)
	}
	index = util.IntIndexOf(1, circle)
	return circle[index+1] + circle[index+2]
}

func print(args ...interface{}) {
	fmt.Println(args...)
}

func parseInput(lines []string) (out []int) {
	l := strings.Split(lines[0], "")
	for _, x := range l {
		out = append(out, util.ParseInt(x))
	}
	return out
}

func completeCircle(in []int) (out []int) {
	for _, x := range in {
		out = append(out, x)
	}
	for i := 10; i <= 1000000; i++ {
		out = append(out, i)
	}
	return out
}

func formatCircle(in []int) (out string) {
	index := util.IntIndexOf(1, in)
	for i := index; i < len(in); i++ {
		out += fmt.Sprintf("%d", in[i])
	}
	for i := 0; i < index; i++ {
		out += fmt.Sprintf("%d", in[i])
	}
	return out[1:]
}

func step(in []int, index int) (out []int, outindex int) {
	current := in[index]
	next := in[(index+4)%len(in)]
	remaining, pickedUp := pickUp(in, index)
	destination := getDestination(current, len(in), pickedUp)
	out = insert(remaining, destination, pickedUp)
	//	fmt.Println(in, remaining, pickedUp, destination, out)
	return out, util.IntIndexOf(next, out)
}

func pickUp(in []int, index int) (out, picked []int) {
	index++
	if index >= len(in) {
		index = 0
	}
	if index+2 < len(in) {
		for i := 0; i < index; i++ {
			out = append(out, in[i])
		}
		for i := index; i < index+3; i++ {
			picked = append(picked, in[i])
		}
		for i := index + 3; i < len(in); i++ {
			out = append(out, in[i])
		}
	} else {
		for i := index; i < index+3; i++ {
			picked = append(picked, in[i%len(in)])
		}
		for i := (index + 3) % len(in); i < index; i++ {
			out = append(out, in[i])
		}
	}

	return out, picked
}

func getDestination(current, max int, in []int) int {
	for {
		current--
		if current == 0 {
			current = max
		}
		if !util.IntInSlice(current, in) {
			return current
		}
	}
}

func getDestinationLinked(current, max int, pickedUp, circle *cup) int {
LOOP:
	for {
		current--
		if current == 0 {
			current = max
		}
		tmp := pickedUp
		for i := 0; i < 3; i++ {
			if tmp.value == current {
				continue LOOP
			}
			tmp = tmp.next
		}
		return current
	}
}

func insertLinked(circle, pickedUp, destination *cup) {
	next := destination.next
	destination.next = pickedUp
	for i := 0; i < 2; i++ {
		pickedUp = pickedUp.next
	}
	pickedUp.next = next
}

func insert(in []int, destination int, picked []int) (out []int) {
	index := util.IntIndexOf(destination, in) + 1
	for i := 0; i < index; i++ {
		out = append(out, in[i])
	}
	for _, x := range picked {
		out = append(out, x)
	}
	for i := index; i < len(in); i++ {
		out = append(out, in[i])
	}
	return out
}

func printLinked(circle *cup, max int) {
	head := circle
	for i := 0; i < max; i++ {
		fmt.Printf("%d ", head.value)
		head = head.next
	}
	fmt.Println()
}

func formatLinked(circle *cup) (out string) {
	for circle.value != 1 {
		circle = circle.next
	}
	circle = circle.next
	for circle.value != 1 {
		out += fmt.Sprintf("%d", circle.value)
		circle = circle.next
	}
	return out
}

func pickUpLinked(circle *cup) (picked *cup, out *cup) {
	head := circle
	picked = circle.next
	fourth := circle
	for i := 0; i < 4; i++ {
		fourth = fourth.next
	}
	head.next = fourth
	return picked, fourth
}
