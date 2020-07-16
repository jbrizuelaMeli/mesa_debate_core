package main

import "fmt"

func passByRefAndValue() {
	x := 0
	y := 5
	fmt.Printf("before:\ta=%d\tb=%d\n", x, y)
	updateValue(x, &y)
	fmt.Printf("after:\ta=%d\tb=%d\n", x, y)
}

func updateValue(a int, b *int) {
	a = 10
	*b = *b * *b
	fmt.Printf("inside:\ta=%d\tb=%d\n", a, *b)
}
