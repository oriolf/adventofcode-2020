package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve(f1, f2 func([]string) interface{}) {
	l := ScanLines()
	switch os.Args[1] {
	case "1":
		fmt.Println(f1(l))
	case "2":
		fmt.Println(f2(l))
	default:
		panic("unknown part")
	}
}

func ScanLines() (out []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		out = append(out, strings.TrimSpace(scanner.Text()))
	}
	return out
}

func ParseInt(s string) int {
	x, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(fmt.Sprintf("not a number: %q", s))
	}
	return x
}

func ParseInts(l []string) (out []int) {
	for _, x := range l {
		out = append(out, ParseInt(x))
	}
	return out
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sum(l []int) (sum int) {
	for _, x := range l {
		sum += x
	}
	return sum
}

func IndexOf(s string, l []string) int {
	for i, x := range l {
		if x == s {
			return i
		}
	}
	return -1
}

func IntIndexOf(s int, l []int) int {
	for i, x := range l {
		if x == s {
			return i
		}
	}
	return -1
}

func StringInSlice(s string, l []string) bool {
	return IndexOf(s, l) != -1
}

func IntInSlice(i int, l []int) bool {
	return IntIndexOf(i, l) != -1
}

func IntersectStrings(a, b []string) (out []string) {
	ma := presenceMap(a)
	mb := presenceMap(b)
	for k := range ma {
		if _, ok := mb[k]; ok {
			out = append(out, k)
		}
	}
	return out
}

func presenceMap(a []string) map[string]struct{} {
	m := make(map[string]struct{}, len(a))
	for _, x := range a {
		m[x] = struct{}{}
	}
	return m
}

func GenerateBoolMatrix(x, y int) (out [][]bool) {
	for i := 0; i < y; i++ {
		out = append(out, GenerateBoolRow(x))
	}
	return out
}

func GenerateBoolRow(x int) (out []bool) {
	for j := 0; j < x; j++ {
		out = append(out, false)
	}
	return out
}

func GenerateIntMatrix(x, y int) (out [][]int) {
	for i := 0; i < y; i++ {
		var r []int
		for j := 0; j < x; j++ {
			r = append(r, 0)
		}
		out = append(out, r)
	}
	return out
}
