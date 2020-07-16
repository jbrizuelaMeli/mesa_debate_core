package main

import "fmt"

func escapeAnalysis() {
	a := new(int)
	b := new(int)

	value(*a)
	//reference(a)
	reference(b)

	sum := *a + *b
	fmt.Println("sum = ", sum)
}

func value(a int){
	//En este caso 'a' necesita ser compartida entre los stack frames de dos funciones [main() y Println()]
	fmt.Println("value of a: ", a)
}

func reference(a *int){
	fmt.Println("direction of a: ", a)
}


