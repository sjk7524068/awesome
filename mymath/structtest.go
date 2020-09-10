package mymath

import "fmt"

type AnimalSkill interface {
	speak() string
}

type Animal struct {
	Name string
	age int
}

func (animal Animal) speak() {
	fmt.Println("animal can speak")
}

