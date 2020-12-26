package main

import (
	"github.com/oriolf/adventofcode2020/util"
	"strings"
)

type instruction struct {
	name  string
	value int
}

func main() {
	util.Solve(solve1, solve2)
}

func solve1(lines []string) interface{} {
	program := parseProgram(lines)
	acc, _ := executeProgram(program)
	return acc
}

func solve2(lines []string) interface{} {
	program := parseProgram(lines)

	for index := 0; index < len(program); index++ {
		if acc, ok := executeProgram(modifyProgram(program, index)); ok {
			return acc
		}
	}

	return 0
}

func parseProgram(lines []string) (program []instruction) {
	for _, l := range lines {
		p := strings.Fields(l)
		program = append(program, instruction{name: p[0], value: util.ParseInt(p[1])})
	}

	return program
}

func modifyProgram(in []instruction, index int) (out []instruction) {
	for i, x := range in {
		if index == i {
			if x.name == "nop" {
				x.name = "jmp"
			} else if x.name == "jmp" {
				x.name = "nop"
			}
		}
		out = append(out, x)
	}
	return out
}

func executeProgram(program []instruction) (acc int, ok bool) {
	var index int
	executed := map[int]struct{}{}
LOOP:
	for index != len(program) {
		if _, ok := executed[index]; ok {
			return acc, false
		}

		executed[index] = struct{}{}
		ins := program[index]
		switch ins.name {
		case "nop":
		case "acc":
			acc += ins.value
		case "jmp":
			index += ins.value
			continue LOOP
		default:
			panic("unknown instruction")
		}

		index++
	}

	return acc, true
}
