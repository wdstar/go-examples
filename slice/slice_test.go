package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	assert := assert.New(t)

	var s1 []string         // zero value (= nil), NOT empty
	s2 := []string{}        // empty
	s3 := make([]string, 0) // empty

	assert.NotEqual(s1, s2)
	assert.NotEqual(s1, s3)
	/*
		expected: []string(nil)
		actual  : []string{}
	*/
	assert.Equal(s2, s3)

	assert.Equal(0, len(s1))
	assert.Equal(0, len(s2))
	assert.Equal(0, len(s3))

	assert.Equal(0, len(append(s1, s2...)))
	assert.Equal(0, len(append(s2, s3...)))
	assert.Equal(0, len(append(s1, s3...)))

	s1 = append(s1, "added element")
	assert.Equal(1, len(append(s1, s2...)))
	assert.Equal(1, len(append(s1, nil...)))
	assert.Equal(1, len(append(s2, s1...)))

	s1 = nil // empty
	s2 = nil // empty
	assert.Equal(0, len(append(s1, s2...)))
	assert.Equal(0, len(append(s1, nil...)))
	assert.Equal(0, len(append(s2, s1...)))
}
