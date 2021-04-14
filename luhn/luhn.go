package luhn

import (
	"regexp"
	"strings"
)

func Valid(input string) bool {
	inputWithoutSpaces := strings.ReplaceAll(input, " ", "")
	if len(inputWithoutSpaces) == 1 {
		return false
	}
	stringContainsNonDigit, _ := regexp.MatchString("[^\\d]", inputWithoutSpaces)
	if stringContainsNonDigit {
		return false
	}

	checksum := 0
	remainder := len(inputWithoutSpaces) % 2
	for index, char := range inputWithoutSpaces {
		number := int(char - '0')
		if index%2 == remainder {
			number = number * 2
			if number > 9 {
				number = number - 9
			}
		}
		checksum += number
	}
	return checksum%10 == 0
}
