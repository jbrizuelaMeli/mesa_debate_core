package main

import "fmt"

func sliceVsArraysVsMaps() {
	fmt.Println("Update Array")
	x := [2]int{1, 2}
	fmt.Println(x)
	updateArray(x)
	fmt.Println(x)

	fmt.Println("Update Slice")
	y := []int{1, 2}
	fmt.Println(y)
	updateSlice(y)
	fmt.Println(y)

	fmt.Println("Update Map")
	z := make(map[int]int)
	z[0] = 1
	z[1] = 2
	fmt.Println(z)
	updateMap(z)
	fmt.Println(z)
}

func updateSlice(a []int) {
	for i, _ := range a {
		//fmt.Println("En update: ", idx, &a[i])
		a[i] *= 10
	}
}

func updateMap(a map[int]int) {
	for i, _ := range a {
		//fmt.Println("En update: ", idx, &a[i])
		a[i] *= 10
	}
}

func updateArray(a [2]int) {
	for i, _ := range a {
		//fmt.Println("En update: ", idx, &a[i])
		a[i] *= 10
	}
}
