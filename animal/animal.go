package animal

import (
	"fmt"
	"io"
	"os"
)

var writer io.Writer

func init() {
	// explicit writer for unit test
	writer = os.Stdout
}

// Animal behavior interface
type Animal interface {
	String() string
	Cry()
	DelegateCry()
	CallCry(a Animal)
}

type animal struct {
	name string
}

// implementation check
var _ Animal = (*animal)(nil)

// NewAnimal creates an animal.
func NewAnimal(name string) Animal {
	return &animal{
		name: name,
	}

	// alternative
	//a := new(animal)
	//a.name = name
	//return a
}

func (a *animal) String() string {
	return fmt.Sprintf("This is an animal (%s).", a.name)
}

func (a *animal) Cry() {
	fmt.Fprintln(writer, "cry")
}

func (a *animal) DelegateCry() {
	a.Cry()
}

func (a *animal) CallCry(aa Animal) {
	aa.Cry()
}
