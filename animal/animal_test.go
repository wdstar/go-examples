package animal

import (
	"bytes"
	"os"
	"testing"
)

var testBuf *bytes.Buffer

func init() {
	// These initializaitons is moved to TestMain()
	//testBuf = &bytes.Buffer{}
	//writer = testBuf
}

func TestMain(m *testing.M) {
	// setup
	testBuf = &bytes.Buffer{}
	writer = testBuf
	code := m.Run()
	// teardown
	writer = os.Stdout
	os.Exit(code)
}

func TestAnimalString(t *testing.T) {
	a := NewAnimal("No name")
	expected := "This is a animal (No name)."
	actual := a.String()
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestAnimalCry(t *testing.T) {
	testBuf.Reset()

	a := NewAnimal("No name")
	expected := "cry\n"
	a.Cry()
	actual := testBuf.String()
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
