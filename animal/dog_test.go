package animal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test helpers for each test case.
func createDog() *Dog {
	return NewDog("Hachi")
}

func TestDogCry(t *testing.T) {
	assert := assert.New(t)
	setUp()

	d := createDog()
	expected := "bowwow\n"
	d.Cry()
	actual := testBuf.String()
	assert.Equal(expected, actual)
	//if actual != expected {
	//	t.Errorf("got: %v\nwant: %v", actual, expected)
	//}
}
