package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var max int
	present := map[int]struct{}{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row, column := parseBoardingPass(scanner.Text())
		res := row*8 + column
		present[res] = struct{}{}
		if res > max {
			max = res
		}
	}

	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			res := i*8 + j
			if _, ok := present[res]; !ok {
				fmt.Println("NO", res)
				_, ok1 := present[res+1]
				_, ok2 := present[res-1]
				if ok1 && ok2 {
					fmt.Println(res)
					os.Exit(0)
				}
			}
		}
	}

	//	fmt.Println(max)
}

func binToInt(bin []bool) (out int) {
	mult := 1
	for i := len(bin) - 1; i >= 0; i-- {
		if bin[i] {
			out += mult
		}
		mult *= 2
	}
	return out
}

func parseBoardingPass(s string) (int, int) {
	var row, column []bool
	for i := 0; i < 7; i++ {
		if s[i:i+1] == "B" {
			row = append(row, true)
		} else {
			row = append(row, false)
		}
	}

	for i := 7; i < 10; i++ {
		if s[i:i+1] == "R" {
			column = append(column, true)
		} else {
			column = append(column, false)
		}
	}

	return binToInt(row), binToInt(column)
}
