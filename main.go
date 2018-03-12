package main

import (
	"fmt"

	"github.com/wdstar/go-examples/animal"
)

func main() {
	fmt.Println("Hello World!")

	a := animal.NewAnimal("No name")
	fmt.Println(a)
	a.Cry()
	a.DelegateCry()
	a.CallCry(a)

	c := animal.NewCat("Mike")
	fmt.Println(c)
	c.Cry()
	c.DelegateCry()
	c.CallCry(c)

	d := animal.NewDog("Hachi")
	fmt.Println(d)
	d.Cry()
	d.DelegateCry()
	d.CallCry(d)
}
