package animal

import "fmt"

// Cat is an animal.
type Cat struct {
	*animal
	age int
}

// NewCat creates a Cat struct.
func NewCat(name string) *Cat {
	// init with tags
	c := &Cat{
		animal: &animal{
			name: name,
		},
		age: 0,
	}

	// alternative 0 (without tags)
	//c = &Cat{
	//	&animal{
	//		name: name,
	//	},
	//	0,
	//}

	// alternative 1
	//c := &Cat{
	//	&animal{},
	//}
	//c.name = name
	//c.age = 0

	// alternative 2 (?!)
	//c := new(Cat)
	//c.animal = new(animal)
	//c.name = name
	//c.age = 0

	return c
}

func (c *Cat) String() string {
	return fmt.Sprintf("This is a cat (%s).", c.name)
}

// Cry overrides animal.Cry()
func (c *Cat) Cry() {
	fmt.Fprintln(writer, "mew")
}
