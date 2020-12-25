package main

import (
	"fmt"
	"github.com/oriolf/adventofcode2020/util"
	"strings"
)

func main() {
	util.Solve(solve(update1), solve(update2))
}

func solve(updateFunc func(map[int]int, []string, string) map[int]int) func(lines []string) interface{} {
	return func(lines []string) interface{} {
		mem := make(map[int]int)
		mask := []string{}
		for _, l := range lines {
			if strings.HasPrefix(l, "mask") {
				mask = parseMask(l)
			} else {
				mem = updateFunc(mem, mask, l)
			}
		}

		var sum int
		for _, v := range mem {
			sum += v
		}

		return sum
	}
}

func update1(mem map[int]int, mask []string, l string) map[int]int {
	var k, v int
	if _, err := fmt.Sscanf(l, "mem[%d] = %d", &k, &v); err != nil {
		panic("bad mem assignment")
	}
	mem[k] = applyMask(mask, v)
	return mem
}

func parseMask(s string) (out []string) {
	for _, x := range strings.Split(s, "") {
		out = append([]string{x}, out...)
	}

	return out
}

func applyMask(mask []string, value int) (out int) {
	for i := 0; i < 36; i++ {
		switch mask[i] {
		case "1":
			out += 1 << i
		case "0":
		case "X":
			out += value & (1 << i)
		default:
			panic("wrong mask value")
		}
	}

	return out
}

func update2(mem map[int]int, mask []string, l string) map[int]int {
	var k, v int
	if _, err := fmt.Sscanf(l, "mem[%d] = %d", &k, &v); err != nil {
		panic("bad mem assignment")
	}

	for _, addr := range computeAddresses(k, mask) {
		mem[addr] = v
	}

	return mem
}

func computeAddresses(k int, mask []string) (out []int) {
	out = []int{k}
	for i := 0; i < 36; i++ {
		switch mask[i] {
		case "1":
			for j := range out {
				out[j] |= 1 << i
			}
		case "0":
		case "X":
			var out2 []int
			for _, x := range out {
				out2 = append(out2, x&(^(1 << i)))
				out2 = append(out2, x|1<<i)
			}
			out = out2
		default:
			panic("wrong mask value")
		}
	}

	return out
}
