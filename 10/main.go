package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	numbers := []int{0}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("not a number")
		}
		numbers = append(numbers, x)
	}

	sort.Ints(numbers)
	//	ones, threes := countOnesThrees(numbers)
	//	fmt.Println(ones, threes, ones*threes)

	differences := computeDifferences(numbers)
	mult := 1
	for {
		consecutiveOnes, nextIndex := countConsecutiveOnes(differences)
		//		fmt.Println(differences, consecutiveOnes, nextIndex)
		switch consecutiveOnes {
		case 1:
		case 2:
			mult *= 2
		case 3:
			mult *= 4
		case 4:
			mult *= 7
		default:
			panic(fmt.Sprintf("unexpected consecutive ones %d", consecutiveOnes))
		}
		if nextIndex >= len(differences) {
			break
		}
		differences = differences[nextIndex:]
	}

	fmt.Println(mult)
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
