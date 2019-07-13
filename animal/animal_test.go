package animal

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testBuf *bytes.Buffer

func init() {
	// These initializations are moved to TestMain()
	//testBuf = &bytes.Buffer{}
	//writer = testBuf
}

func TestMain(m *testing.M) {
	// before all
	testBuf = &bytes.Buffer{}
	writer = testBuf

	code := m.Run()

	// after all
	writer = os.Stdout
	os.Exit(code)
}

// test helpers for each test case.
func setUp() {
	testBuf.Reset()
}

func tearDown() {
}

func createAnimal() Animal {
	return NewAnimal("No name")
}

func TestAnimalString(t *testing.T) {
	assert := assert.New(t)

	a := createAnimal()
	expected := "This is an animal (No name)."
	actual := a.String()
	assert.Equal(expected, actual)
	//if actual != expected {
	//	t.Errorf("got: %v\nwant: %v", actual, expected)
	//}
}

func TestAnimalCry(t *testing.T) {
	assert := assert.New(t)
	setUp()

	a := createAnimal()
	expected := "cry\n"
	a.Cry()
	actual := testBuf.String()
	assert.Equal(expected, actual)
	//if actual != expected {
	//	t.Errorf("got: %v\nwant: %v", actual, expected)
	//}
}
