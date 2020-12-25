package main

import (
	"github.com/oriolf/adventofcode2020/util"
	"strconv"
	"strings"
)

func main() {
	util.Solve(solve(validPassport1), solve(validPassport2))
}

func solve(validFunc func(map[string]string) bool) func([]string) interface{} {
	return func(lines []string) interface{} {
		var passports []map[string]string
		var text string
		for _, t := range lines {
			if t == "" {
				passports = append(passports, parsePassport(text))
				text = ""
			} else {
				text += t + " "
			}
		}
		passports = append(passports, parsePassport(text))

		var count int
		for _, x := range passports {
			if validFunc(x) {
				count++
			}
		}

		return count
	}
}

var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

var fieldsValidation = map[string]func(string) bool{
	"byr": validateNumberBetween(1920, 2002),
	"iyr": validateNumberBetween(2010, 2020),
	"eyr": validateNumberBetween(2020, 2030),
	"hgt": func(s string) bool {
		if !strings.HasSuffix(s, "cm") && !strings.HasSuffix(s, "in") {
			return false
		}

		min, max := 150, 193
		if s[len(s)-2:] == "in" {
			min, max = 59, 76
		}
		s = s[:len(s)-2]
		x, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return x >= min && x <= max
	},
	"hcl": func(s string) bool {
		if len(s) != 7 {
			return false
		}
		if s[0:1] != "#" {
			return false
		}
		for i := 1; i < 7; i++ {
			if !util.StringInSlice(s[i:i+1], []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}) {
				return false
			}
		}
		return true
	},
	"ecl": func(s string) bool {
		return util.StringInSlice(s, []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"})
	},
	"pid": func(s string) bool {
		if len(s) != 9 {
			return false
		}
		_, err := strconv.Atoi(s)
		return err == nil
	},
}

func validateNumberBetween(min, max int) func(string) bool {
	return func(s string) bool {
		if len(s) != 4 {
			return false
		}
		x, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return x >= min && x <= max
	}
}

func parsePassport(s string) map[string]string {
	fields := strings.Fields(s)
	m := make(map[string]string)
	for _, x := range fields {
		parts := strings.Split(x, ":")
		m[parts[0]] = parts[1]
	}

	return m
}

func validPassport1(m map[string]string) bool {
	for _, f := range requiredFields {
		if _, ok := m[f]; !ok {
			return false
		}
	}
	return true
}

func validPassport2(m map[string]string) bool {
	for _, f := range requiredFields {
		if v, ok := m[f]; !ok {
			return false
		} else {
			if fun := fieldsValidation[f]; !fun(v) {
				return false
			}
		}
	}
	return true
}
