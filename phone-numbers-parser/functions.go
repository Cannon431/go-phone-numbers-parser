package phone_numbers_parser

import (
	"os"
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

func filter(number string) (bool, error) {
	number = strings.TrimSpace(number)
	correct := true

	minDigitsCount, err := strconv.Atoi(os.Getenv("min_digits_count"))

	if err != nil {
		return false, err
	}

	if countOfDigits(number) < minDigitsCount {
		correct = false
	}

	ignored := strings.Split(os.Getenv("ignored"), ",")

	isInIgnored := false

	for _, filter := range ignored {
		if number == filter {
			isInIgnored = true
		}
	}

	if isInIgnored {
		correct = false
	}

	return correct, nil
}

func mustFilter(number string) bool {
	correct, err := filter(number)

	if err != nil {
		panic(err)
	}

	return correct
}
