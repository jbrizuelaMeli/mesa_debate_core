package main

import "fmt"

type HelloFunc func(string)

func SayHello(to string) {
	fmt.Printf("Hello, %s!\n", to)
}

func functionParam2() {
	var hf HelloFunc

	hf = SayHello
	hf2 := SayHello

	hf("world")
	hf2("world2")

	println(SayHello)
	println(hf)
	println(hf2)
}
