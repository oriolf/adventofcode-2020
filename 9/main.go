package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const preamble = 25

//const target = 127
const target = 675280050

func main() {
	var numbers []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("not a number")
		}
		numbers = append(numbers, x)
	}

	// first part
	//	for i := preamble; i < len(numbers); i++ {
	//		x, y := sumPair(numbers[i-preamble:i], numbers[i])
	//		if x == 0 && y == 0 {
	//			fmt.Println(numbers[i])
	//			os.Exit(0)
	//		}
	//	}

	for i := 0; i < len(numbers); i++ {
		slice, ok := findSliceSum(numbers[i:], target)
		if ok {
			fmt.Println(min(slice) + max(slice))
			os.Exit(0)
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

func findSliceSum(numbers []int, target int) (out []int, ok bool) {
	var sum int
	for _, x := range numbers {
		sum += x
		out = append(out, x)
		if sum == target && len(out) > 1 {
			return out, true
		} else if sum > target {
			return nil, false
		}
	}

	return nil, false
}

func min(l []int) (m int) {
	m = l[0]
	for _, x := range l {
		if x < m {
			m = x
		}
	}
	return m
}

func max(l []int) (m int) {
	for _, x := range l {
		if x > m {
			m = x
		}
	}
	return m
}
