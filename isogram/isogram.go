package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(input string) bool {
	lowercaseInput := strings.ToLower(input)

	for index, letterToTest := range lowercaseInput {
		if !unicode.IsLetter(letterToTest) {
			continue
		}
		for _, letterInRemainingString := range lowercaseInput[index+1:] {
			if letterToTest == letterInRemainingString {
				return false
			}
		}
	}
	return true
}
