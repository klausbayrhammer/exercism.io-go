package grains

import "errors"

func Square(input int) (uint64, error) {
	numberOfSquare := int(input)
	if numberOfSquare <= 0 || numberOfSquare > 64 {
		return 0, errors.New("Invalid input")
	}
	grains := uint64(1)
	for i := 1; i < numberOfSquare; i++ {
		grains *= 2
	}
	return grains, nil
}

func Total() uint64 {
	return 18446744073709551615
}
