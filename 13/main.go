package main

import (
	"github.com/oriolf/adventofcode2020/util"
	"strings"
)

type bus struct {
	value, offset int
}

func main() {
	util.Solve(solve1, solve2)
}

func solve1(lines []string) interface{} {
	t := util.ParseInt(lines[0])
	buses := parseBuses(lines[1])
	min, minBus := int(1e10), -1
	for _, bus := range buses {
		x := bus.value - (t % bus.value)
		if x < min {
			min = x
			minBus = bus.value
		}
	}

	return min * minBus
}

func solve2(lines []string) interface{} {
	buses := parseBuses(lines[1])

	minimum := buses[0].value
	delta := minimum
	for _, b := range buses[1:] {
		minimum, delta = findLowest(b, minimum, delta)
	}

	return minimum
}

func parseBuses(l string) (buses []bus) {
	for i, x := range strings.Split(l, ",") {
		if x == "x" {
			continue
		}
		buses = append(buses, bus{value: util.ParseInt(x), offset: i})
	}
	return buses
}

func findLowest(b bus, minimum, delta int) (int, int) {
	for i := minimum; true; i += delta {
		if (i+b.offset)%b.value == 0 {
			return i, delta * b.value
		}
	}

	return -1, -1
}
