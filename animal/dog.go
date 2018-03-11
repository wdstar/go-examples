package animal

import "fmt"

// Dog is an animal.
type Dog struct {
	*animal
}

// NewDog creates a dog.
func NewDog(name string) *Dog {
	return &Dog{
		&animal{
			name: name,
		},
	}
}

// Cry overrides animal.Cry()
func (d *Dog) Cry() {
	fmt.Fprintln(writer, "bowwow")
}
