package animal

import (
	"testing"
)

// test helpers for each test case.
func createCat() *Cat {
	return NewCat("Mike")
}

func TestCatString(t *testing.T) {
	c := createCat()
	expected := "This is a cat (Mike)."
	actual := c.String()
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func TestCatCry(t *testing.T) {
	setUp()

	a := createCat()
	expected := "mew\n"
	a.Cry()
	actual := testBuf.String()
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
