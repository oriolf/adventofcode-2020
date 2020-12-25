package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	name  string
	value int
}

func main() {
	var program []instruction
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		program = append(program, parseInstruction(text))
	}

	for index := 0; index < len(program); index++ {
		if acc, ok := executeProgram(modifyProgram(program, index)); ok {
			fmt.Println(acc)
			os.Exit(0)
		}
	}

	panic("no index works!")
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

func parseInstruction(s string) instruction {
	parts := strings.Fields(s)
	if len(parts) != 2 {
		panic("wrong instruction")
	}
	x, err := strconv.Atoi(parts[1])
	if err != nil {
		panic("wrong instruction value")
	}
	return instruction{name: parts[0], value: x}
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
