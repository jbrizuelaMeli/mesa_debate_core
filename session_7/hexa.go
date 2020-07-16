package main

import "fmt"

func hexa() {

	// storing the hexadecimal
	// values in variables
	x := 0xFF
	y := 0x9C

	// Displaying the values
	fmt.Printf("El tipo de la variable x es %T\n", x)
	fmt.Printf("El valor de x en hexadecimal es %X\n", x)
	fmt.Printf("El valor de x en decimal es %v\n", x)
	fmt.Printf("La dirección de memoria de x is %v\n", &x)

	fmt.Printf("El tipo de la variable y es %T\n", y)
	fmt.Printf("El valor de y en hexadecimal es %X\n", y)
	fmt.Printf("El valor de y en decimal es %v\n", y)
	fmt.Printf("La dirección de memoria de y is %v\n", &y)
}
