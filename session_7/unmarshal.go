package main

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	X int  `json:"x"`
	Y int  `json:"y"`
	Z *int `json:"z"`
}

type Tank struct {
	Name     string   `json:"name"`
	Position Point    `json:"position"`
	Drivers  []string `json:"drivers"`
}

func unmarshal() {
	myTank := Tank{
		Name: "tank_1234",
		Position: Point{
			X: 12,
			Y: 34,
		},
		Drivers: []string{"man_1", "man_2", "man_3"},
	}
	jsonData, _ := json.Marshal(&myTank)
	fmt.Println(string(jsonData)) // json data

	var newTank Tank
	_ = json.Unmarshal(jsonData, &newTank)
	fmt.Println(newTank) // tank data
}
