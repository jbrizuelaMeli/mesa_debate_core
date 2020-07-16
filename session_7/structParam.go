package main

import "fmt"

type Person struct {
	name     *string
	age      *int
	lastName string
}

func structParams() {
	name := "juan"
	age := 14
	lastName := "Brizuela"

	person := Person{
		name:     &name,
		age:      &age,
		lastName: lastName,
	}

	fmt.Println("En main: ", person, &person.lastName)
	fmt.Println(*person.age, *person.name, person.lastName)
	update(person)
	fmt.Println(age, name, lastName)
}

func update(a Person) {
	fmt.Println("En update: ", a, &a.lastName)
	name := "Pedro"
	age := 21
	*(a.name) = name
	*(a.age) = age
	a.lastName = "Bielli"
}