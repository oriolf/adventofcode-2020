package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} //, "cid"}

var fieldsValidation = map[string]func(string) bool{
	"byr": func(s string) bool {
		if len(s) != 4 {
			return false
		}
		x, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return x >= 1920 && x <= 2002
	},
	"iyr": func(s string) bool {
		if len(s) != 4 {
			return false
		}
		x, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return x >= 2010 && x <= 2020
	},
	"eyr": func(s string) bool {
		if len(s) != 4 {
			return false
		}
		x, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return x >= 2020 && x <= 2030
	},
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
			if !stringInSlice(s[i:i+1], []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}) {
				return false
			}
		}
		return true
	},
	"ecl": func(s string) bool {
		return stringInSlice(s, []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"})
	},
	"pid": func(s string) bool {
		if len(s) != 9 {
			return false
		}
		_, err := strconv.Atoi(s)
		return err == nil
	},
}

func stringInSlice(s string, l []string) bool {
	for _, x := range l {
		if x == s {
			return true
		}
	}
	return false
}

func main() {
	var passports []map[string]string
	var text string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if t := scanner.Text(); strings.TrimSpace(t) == "" {
			passports = append(passports, parsePassport(text))
			text = ""
		} else {
			text += t + " "
		}
	}
	passports = append(passports, parsePassport(text))

	var count int
	for _, x := range passports {
		if validPassport2(x) {
			count++
		}
	}

	fmt.Println(count)
}

func parsePassport(s string) map[string]string {
	fields := strings.Fields(s)
	m := make(map[string]string)
	for _, x := range fields {
		parts := strings.Split(x, ":")
		if len(parts) != 2 {
			panic("wrong parts length")
		}
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
			fun := fieldsValidation[f]
			if !fun(v) {
				return false
			}
		}
	}
	return true
}
