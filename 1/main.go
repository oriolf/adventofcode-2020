package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var numbers []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if n, err := strconv.Atoi(scanner.Text()); err != nil {
			panic("non int input")
		} else {
			numbers = append(numbers, n)
		}
	}

	// 	fmt.Println(sumPair(numbers, 2020)) // first problem
	for i, n := range numbers[1:] {
		x, y := sumPair(numbers[:i], 2020-n)
		if x+y == 2020-n {
			fmt.Println(x * y * n)
		}
	}
}

func sumPair(numbers []int, target int) (int, int) {
	seen := map[int]struct{}{}
	for _, n := range numbers {
		if _, ok := seen[target-n]; ok {
			return n, target - n
		} else {
			seen[n] = struct{}{}
		}
	}
	return 0, 0
}
