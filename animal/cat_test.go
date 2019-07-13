package animal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test helpers for each test case.
func createCat() *Cat {
	return NewCat("Mike")
}

func TestCatString(t *testing.T) {
	assert := assert.New(t)

	c := createCat()
	expected := "This is a cat (Mike)."
	actual := c.String()
	assert.Equal(expected, actual)
	//if actual != expected {
	//	t.Errorf("got: %v\nwant: %v", actual, expected)
	//}
}

func TestCatCry(t *testing.T) {
	assert := assert.New(t)
	setUp()

	a := createCat()
	expected := "mew\n"
	a.Cry()
	actual := testBuf.String()
	assert.Equal(expected, actual)
	//if actual != expected {
	//	t.Errorf("got: %v\nwant: %v", actual, expected)
	//}
}
