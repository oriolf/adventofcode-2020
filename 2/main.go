package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type testcase struct {
	min, max    int
	char, value string
}

func main() {
	var cases []testcase
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cases = append(cases, getCase(scanner.Text()))
	}

	var count int
	for _, x := range cases {
		if isOk2(x) {
			count++
		}
	}

	fmt.Println(count)
}

func getCase(s string) (out testcase) {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, ":", "")
	if _, err := fmt.Sscanf(s, "%d %d %s %s", &out.min, &out.max, &out.char, &out.value); err != nil {
		panic(err)
	}
	return out
}

func isOk1(x testcase) bool {
	var count int
	for i := range x.value {
		if x.value[i:i+1] == x.char {
			count++
		}
	}
	return count >= x.min && count <= x.max
}

func isOk2(x testcase) bool {
	var count int
	for i := range x.value {
		if (i == x.min-1 || i == x.max-1) && x.value[i:i+1] == x.char {
			count++
		}
	}
	return count == 1
}
