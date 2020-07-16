package main

import "fmt"

func pointerOperations() {
	var x = 67
	//var p = &x
	//var p1 = p + 1 // Compiler Error: invalid operation

	var p1 = &x
	var p2 = &x

	if p1 == p2 {
		fmt.Println("Both pointers p1 and p2 point to the same variable.")
	}

	/*if p1 > p2 {
		fmt.Println("Both pointers p1 and p2 point to the same variable.")
	}*/
}
