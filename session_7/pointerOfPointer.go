package main

import "fmt"

func pointerOfPointer() {
	var a = 7.98
	var p = &a
	var pp = p

	fmt.Println("a = ", a)
	fmt.Println("address of a = ", &a)

	fmt.Println("p = ", p)
	fmt.Println("address of p = ", &p)
	fmt.Println("address of pp = ", &pp)

	fmt.Println("pp = ", pp)

	fmt.Println("*pp = ", *pp)
}
