package main

import "fmt"

func closure() {
	nextInt := intSeq()
	fmt.Println("nextInt", nextInt)
	fmt.Printf("nextInt es de tipo %T\n",nextInt)
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt2 := intSeq()
	fmt.Println("newInt2", newInt2)
	fmt.Println(newInt2())
}

func intSeq() func() int {
	i := 0
	fmt.Println("Dirección de i en intSeq: ", &i)
	return func() int {
		fmt.Println("Dirección i: ", &i)
		i += 1
		return i
	}
}
