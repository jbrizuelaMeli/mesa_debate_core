package main

import "fmt"

func pointersAndOperators() {
	var p *int
	i := 42
	p = &i
	fmt.Println(*p) // Obtiene el valor de i a través de p
	*p = 21         // Setea el valor de i a través de p
	fmt.Println(i)  // Obtendría el nuevo valor 21
	fmt.Println(&i) // Imprime la dirección de memoria de i
	fmt.Printf("El tipo de dato de p es %T", p)
}
