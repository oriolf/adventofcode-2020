package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	group := map[string]int{}
	groupTotal := 0
	var groups []map[string]int
	var groupTotals []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			groups = append(groups, group)
			groupTotals = append(groupTotals, groupTotal)
			group = map[string]int{}
			groupTotal = 0
		} else {
			groupTotal++
			group = updateGroup(group, text)
		}
	}

	groups = append(groups, group)
	groupTotals = append(groupTotals, groupTotal)

	var sum int
	//	for _, group := range groups {
	//		sum += len(group)
	//	}

	for i, group := range groups {
		for _, v := range group {
			if v == groupTotals[i] {
				sum++
			}
		}
	}

	fmt.Println(sum)
}

func updateGroup(group map[string]int, text string) map[string]int {
	for i := range text {
		group[text[i:i+1]]++
	}
	return group
}
