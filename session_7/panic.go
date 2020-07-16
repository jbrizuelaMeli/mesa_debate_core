package main

import "fmt"

type point struct {
	X int
	Y int
}

func panic() { //Panic runtime
	var p point // non-pointer
	p.X = 1
	p.Y = 2
	var q *point // pointer
	//q = &Point{} // assign before using
	//q.X = 3
	//q.Y = 4
	fmt.Println(p, q)
}
