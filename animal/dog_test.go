package animal

import (
	"testing"
)

func TestDogCry(t *testing.T) {
	testBuf.Reset()

	d := NewDog("Hachi")
	expected := "bowwow\n"
	d.Cry()
	actual := testBuf.String()
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
