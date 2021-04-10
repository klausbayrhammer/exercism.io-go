package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strings dont have equal length")
	}

	distance := 0
	for position := range a {
		if a[position] != b[position] {
			distance++
		}

	}
	return distance, nil
}
