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

	d := animal.NewDog("Hachi")
	fmt.Println(d)
	d.Cry()

	c := animal.NewCat("Mike")
	fmt.Println(c)
	c.Cry()
}
