package base62

import (
	"errors"
	"math"
	"strings"
)

const (
	base         uint64 = 62
	characterSet        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func Encode(num uint64) string {

	encoded := ""

	for num > 0 {
		reminder := num % base
		num /= base
		encoded = string(characterSet[reminder]) + encoded
	}

	return encoded
}

func Decode(encoded string) (uint64, error) {

	var res uint64

	for index, char := range encoded {

		pow := len(encoded) - (index + 1)
		pos := strings.IndexRune(characterSet, char)

		if pos == -1 {
			return 0, errors.New("(ERR) >> Invalid character: " + string(char))
		}

		res += uint64(pos) * uint64(math.Pow(float64(base), float64(pow)))
	}

	return res, nil
}
