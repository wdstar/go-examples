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
	c := createCat()
	expected := "This is a cat (Mike)."
	actual := c.String()
	assert.Equal(t, expected, actual)
	//if actual != expected {
	//	t.Errorf("got: %v\nwant: %v", actual, expected)
	//}
}

func TestCatCry(t *testing.T) {
	setUp()

	a := createCat()
	expected := "mew\n"
	a.Cry()
	actual := testBuf.String()
	assert.Equal(t, expected, actual)
	//if actual != expected {
	//	t.Errorf("got: %v\nwant: %v", actual, expected)
	//}
}
