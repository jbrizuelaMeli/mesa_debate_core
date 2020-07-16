package main

import "fmt"

func functionParam() {
	calculate(Plus)
	fmt.Println("Plus-out: ",Plus)
	calculate(Minus)
	fmt.Println("Minus-out: ",Minus)
	calculate(Multiply)
	fmt.Println("Multiply-out: ",Multiply)
}

func calculate(fp func(int, int) int) {
	ans := fp(3, 2)
	fmt.Printf("\n%v\n", ans)
	fmt.Println("In calculate: ",fp)
}

func Plus(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}
