package base62

import (
	"errors"
	"strings"
)

const (
	base         uint64 = 62
	characterSet        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func Encode(num uint64) string {

	if num == 0 {
		return "0"
	}

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

	for _, char := range encoded {

		pos := strings.IndexRune(characterSet, char)

		if pos == -1 {
			return 0, errors.New("(ERR) >> Invalid character: " + string(char))
		}

		res = res*base + uint64(pos)
	}

	return res, nil
}
