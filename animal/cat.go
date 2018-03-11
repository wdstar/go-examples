package animal

import "fmt"

// Cat is an animal.
type Cat struct {
	*animal
}

// NewCat creates a Cat struct.
func NewCat(name string) *Cat {
	c := &Cat{
		&animal{
			name: name,
		},
	}

	// alternative 1
	//c := &Cat{
	//	&animal{},
	//}
	//c.name = name

	// alternative 2 (?!)
	//c := new(Cat)
	//c.animal = new(animal)
	//c.name = name

	return c
}

func (c *Cat) String() string {
	return fmt.Sprintf("This is a cat (%s).", c.name)
}

// Cry overrides animal.Cry()
func (c *Cat) Cry() {
	fmt.Fprintln(writer, "mew")
}
