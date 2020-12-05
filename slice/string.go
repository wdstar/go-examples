package slice

import (
	"unicode/utf8"
)

func sanitizeInvalidUTF8(target []byte) []byte {
	if utf8.Valid(target) {
		return target
	}

	buf := []byte("")
	for _, runeVal := range string(target) {
		// rune sanitizes invalid utf-8.
		buf = append(buf, string(runeVal)...)
	}

	return buf
}

func sanitizeInvalidUTF8viaRune(target []byte) []byte {
	if utf8.Valid(target) {
		return target
	}

	// sanitize invalid utf-8 via []rune().
	return []byte(string([]rune(string(target))))
}
