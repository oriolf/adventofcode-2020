package main

import (
	"github.com/oriolf/adventofcode2020/util"
	"sort"
)

func main() {
	util.Solve(solve1, solve2)
}

func solve1(lines []string) interface{} {
	numbers := append([]int{0}, util.ParseInts(lines)...)
	sort.Ints(numbers)
	ones, threes := countOnesThrees(numbers)
	return ones * threes
}

func solve2(lines []string) interface{} {
	numbers := append([]int{0}, util.ParseInts(lines)...)
	sort.Ints(numbers)
	differences := computeDifferences(numbers)
	mult := 1
	for {
		consecutiveOnes, nextIndex := countConsecutiveOnes(differences)
		switch consecutiveOnes {
		case 1:
		case 2:
			mult *= 2
		case 3:
			mult *= 4
		case 4:
			mult *= 7
		default:
			panic("unexpected consecutive ones")
		}
		if nextIndex >= len(differences) {
			break
		}
		differences = differences[nextIndex:]
	}

	return mult
}

func computeDifferences(in []int) (out []int) {
	for i := 0; i < len(in)-1; i++ {
		out = append(out, in[i+1]-in[i])
	}
	return out
}

func countConsecutiveOnes(numbers []int) (count, index int) {
	for index < len(numbers) && numbers[index] == 1 {
		count++
		index++
	}
	for index < len(numbers) && numbers[index] != 1 {
		index++
	}
	return count, index
}

func countOnesThrees(numbers []int) (ones, threes int) {
	for i := 0; i < len(numbers)-1; i++ {
		diff := numbers[i+1] - numbers[i]
		if diff == 1 {
			ones++
		} else if diff == 3 {
			threes++
		} else {
			panic("unexpected difference")
		}
	}

	threes++
	return ones, threes
}
