package slice

import (
	"unicode/utf8"
)

func cleanBytes(target []byte) []byte {
	if utf8.Valid(target) {
		return target
	}

	buf := []byte("")
	for _, runeVal := range string(target) {
		buf = append(buf, string(runeVal)...)
	}

	return buf
}
