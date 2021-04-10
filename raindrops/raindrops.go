package raindrops

import "strconv"

func Convert(a int) string {
	answer := ""
	if a%3 == 0 {
		answer += "Pling"
	}
	if a%5 == 0 {
		answer += "Plang"
	}
	if a%7 == 0 {
		answer += "Plong"
	}
	if answer != "" {
		return answer
	}
	return strconv.Itoa(a)
}
