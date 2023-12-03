package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var digitsMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

func One() int {
	b, _ := os.ReadFile("day1/input1.txt")

	s := strings.Split(string(b), "\n")

	var getFirstDigit = func(s string) string {
		for i, r := range s {
			if r >= '0' && r <= '9' {
				return s[i : i+1]
			}
		}
		return ""
	}

	var getLastDigit = func(s string) string {
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] >= '0' && s[i] <= '9' {
				// result, _ := strconv.Atoi(s[i : i+1])
				return s[i : i+1]
			}
		}
		return ""
	}

	res := 0

	for _, v := range s {
		firstDigit := getFirstDigit(v)
		lastDigit := getLastDigit(v)
		number, _ := strconv.Atoi(fmt.Sprintf("%s%s", firstDigit, lastDigit))
		res += number
	}

	return res

}

func Two() int {
	b, _ := os.ReadFile("day1/input1.txt")

	s := strings.Split(string(b), "\n")

	var getFirstDigit = func(s string) string {
		digitValue := ""
		idx := len(s)
		for k, d := range digitsMap {
			var newIdx = strings.Index(s, k)
			if newIdx != -1 && newIdx < idx {
				idx = newIdx
				digitValue = d
			}
		}
		return digitValue
	}

	var getLastDigit = func(s string) string {
		digitValue := ""
		idx := -1
		for k, d := range digitsMap {
			var newIdx = strings.LastIndex(s, k)
			if newIdx != -1 && newIdx > idx {
				idx = newIdx
				digitValue = d
			}
		}
		return digitValue
	}

	res := 0

	for _, v := range s {
		firstDigit := getFirstDigit(v)
		lastDigit := getLastDigit(v)
		number, _ := strconv.Atoi(firstDigit + lastDigit)
		res += number
	}

	return res

}
