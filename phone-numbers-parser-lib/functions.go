package pnplib

import (
	"strconv"
	"strings"
	"unicode"
)

func joinInts(ints []int, delimiter string) string {
	var strs []string

	for _, v := range ints {
		strs = append(strs, strconv.Itoa(v))
	}

	return strings.Join(strs, delimiter)
}

func countOfDigits(str string) int {
	count := 0

	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}

	return count
}
