package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var forest [][]bool
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		forest = append(forest, parseLine(strings.TrimSpace(scanner.Text())))
	}

	// 	fmt.Println(countDescent(forest, 3, 1)) // first problem

	result := 1
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	for _, x := range slopes {
		result *= countDescent(forest, x[0], x[1])
	}
	fmt.Println(result)
}

func parseLine(s string) (out []bool) {
	for i := range s {
		if s[i:i+1] == "#" {
			out = append(out, true)
		} else {
			out = append(out, false)
		}
	}

	return out
}

func countDescent(forest [][]bool, right, down int) (count int) {
	coln := len(forest[0])
	row, column := down, right
	for row < len(forest) {
		if forest[row][column%coln] {
			count++
		}
		row += down
		column += right
	}
	return count
}
