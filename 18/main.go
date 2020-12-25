package main

import (
	"github.com/oriolf/adventofcode2020/util"
	"strings"
	"unicode"
)

type expression struct {
	expressions []expression
	value       int
	op          string
}

func main() {
	util.Solve(solve1, solve2)
}

func solve1(lines []string) interface{} {
	var results []int
	for _, l := range lines {
		result := evaluateExpression(l)
		results = append(results, result)
	}
	return util.Sum(results)
}

func evaluateExpression(s string) int {
	l := parseExpression(s, nil)
	return evaluate(l)
}

func solve2(lines []string) interface{} {
	var results []int
	for _, l := range lines {
		result := evaluateExpression2(l)
		results = append(results, result)
	}
	return util.Sum(results)
}

func evaluateExpression2(s string) int {
	l := parseExpression(s, nil)
	return evaluate2(l)
}

func parseExpression(s string, expressions []expression) []expression {
	s = strings.TrimSpace(s)
	if s == "" {
		return expressions
	}
	if s[0:1] == "(" {
		i := findParenthesisPair(s)
		exprs := parseExpression(s[1:i], nil)
		var op string
		for ; i < len(s); i++ {
			o := s[i : i+1]
			if o == "+" || o == "*" {
				op = o
				i++
				break
			}
		}

		expressions = append(expressions, expression{expressions: exprs, op: op})
		return append(expressions, parseExpression(s[i:], nil)...)
	}

	var i int
	var number string
	for i < len(s) {
		if !unicode.IsDigit(rune(s[i])) {
			break
		}
		number += s[i : i+1]
		i++
	}

	var op string
	for ; i < len(s); i++ {
		o := s[i : i+1]
		if o == "+" || o == "*" {
			op = o
		}

		if o == ")" {
			expressions = append(expressions, expression{value: util.ParseInt(number), op: op})
			return append(expressions, parseExpression(s[i+1:], nil)...)
		}

		if unicode.IsDigit(rune(o[0])) || o == "(" {
			expressions = append(expressions, expression{value: util.ParseInt(number), op: op})
			return append(expressions, parseExpression(s[i:], nil)...)
		}
	}

	return append(expressions, expression{value: util.ParseInt(number)})
}

func evaluate(list []expression) int {
	val := evaluateSingle(list[0])

	for i, e := range list[:len(list)-1] {
		if e.op == "+" {
			val += evaluateSingle(list[i+1])
		} else if e.op == "*" {
			val *= evaluateSingle(list[i+1])
		} else {
			break
		}
	}

	return val
}

func evaluateSingle(e expression) int {
	if len(e.expressions) == 0 {
		return e.value
	}
	return evaluate(e.expressions)
}

func evaluateSingle2(e expression) int {
	if len(e.expressions) == 0 {
		return e.value
	}
	return evaluate2(e.expressions)
}

func findParenthesisPair(s string) int {
	var count int
	for i := range s {
		if s[i:i+1] == "(" {
			count++
		} else if s[i:i+1] == ")" {
			count--
		}
		if count == 0 {
			return i + 1
		}
	}

	return 0
}

func evaluate2(list []expression) int {
	for containsSum(list) {
		list = computeOneSum(list)
	}

	val := evaluateSingle2(list[0])

	for i, e := range list[:len(list)-1] {
		if e.op == "*" {
			val *= evaluateSingle2(list[i+1])
		} else {
			break
		}
	}

	return val
}

func containsSum(list []expression) bool {
	for _, e := range list {
		if e.op == "+" {
			return true
		}
	}
	return false
}

func computeOneSum(list []expression) (out []expression) {
	i := 0
	for _, e := range list {
		if e.op == "+" {
			val := evaluateSingle2(e) + evaluateSingle2(list[i+1])
			out = append(out, expression{value: val, op: list[i+1].op})
			i += 2
			break
		} else {
			out = append(out, e)
		}
		i++
	}

	for ; i < len(list); i++ {
		out = append(out, list[i])
	}

	return out
}
