package animal

import (
	"testing"
)

func TestCatString(t *testing.T) {
	c := NewCat("Mike")
	expected := "This is a cat (Mike)."
	actual := c.String()
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCatCry(t *testing.T) {
	testBuf.Reset()

	a := NewCat("Mike")
	expected := "mew\n"
	a.Cry()
	actual := testBuf.String()
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
