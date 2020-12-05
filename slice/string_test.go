package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeInvalidUTF8(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		target   string
		expected string
	}{
		{"abc", "abc"},
		// invalid utf-8
		{
			"abc\x80", // []byte{0x61, 0x62, 0x63, 0x80}
			"abc�",    // []byte{0x61, 0x62, 0x63, 0xef, 0xbf, 0xbd}
		},
		{
			"\x80abc",
			"�abc",
		},
		{
			"ab\x80c",
			"ab�c",
		},
	}

	for _, c := range cases {
		if c.target != c.expected {
			// []byte() and string() retain invalid utf-8.
			assert.NotEqual(c.expected, string([]byte(c.target)))
		}
		assert.Equal(c.expected, string(sanitizeInvalidUTF8([]byte(c.target))))
		assert.Equal(c.expected, string(sanitizeInvalidUTF8viaRune([]byte(c.target))))
	}
}
