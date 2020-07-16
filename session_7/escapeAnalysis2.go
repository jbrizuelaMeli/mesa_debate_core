package main

import "fmt"

func escapeAnalysis2() {
	//fmt.Println("Call to heapAnalysis", heapAnalysis())
	fmt.Println("Call to stackAnalysis", stackAnalysis())
}

func heapAnalysis() *int {
	data := 55
	return &data
}

func stackAnalysis() int {
	data := 55
	return data
}
