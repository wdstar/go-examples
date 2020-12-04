package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanBytes(t *testing.T) {
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
		assert.Equal(c.expected, string(cleanBytes([]byte(c.target))))
	}
}
