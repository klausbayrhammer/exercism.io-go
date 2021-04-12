package isogram

import "strings"

func IsIsogram(input string) bool {
	lowercaseInput := strings.ToLower(input)
	ignoreCharsReplacer := strings.NewReplacer("-", "", " ", "")
	inputWithoutSpecialChars := ignoreCharsReplacer.Replace(lowercaseInput)

	for index, letterToTest := range inputWithoutSpecialChars {
		for _, letterInRemainingString := range inputWithoutSpecialChars[index+1:] {
			if letterToTest == letterInRemainingString {
				return false
			}
		}
	}
	return true
}
