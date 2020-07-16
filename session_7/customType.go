package main

import "fmt"

type Creature struct {
	Species string
}

func (c *Creature) ResetByReference() {
	c.Species = ""
}

func (c Creature) ResetByValue() {
	c.Species = ""
}

func customType() {
	var creature Creature = Creature{Species: "shark"}

	fmt.Printf("1 Inicio) %+v\n", creature)
	creature.ResetByValue()
	fmt.Printf("2 Despues de pasar por valor) %+v\n", creature)

	fmt.Printf("1 Inicio) %+v\n", creature)
	creature.ResetByReference()
	fmt.Printf("2 Despues de pasar por referencia) %+v\n", creature)
}
