package main

import "fmt"

func pointersDeclaration() {
	var a *int
	var b *string
	var c *float64
	var d *bool
	xPtr := new(int)
	fmt.Println(a, b, c, d, xPtr, *xPtr)
}
