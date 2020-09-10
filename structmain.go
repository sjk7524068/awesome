package main

import "fmt"

type AnimalSkill interface {
	animalspeak() string

}

type LivingSkill interface {
	breath() string
}


type Animal struct {
	Name string
	Age int
	SpeakStyle string
}

func (animal Animal) animalspeak() string{
	fmt.Println(animal.Name, "can speak like", animal.SpeakStyle)
	saying := animal.Name+"can speak like"+ animal.SpeakStyle
	return saying
}

func(animal Animal) breath() string{
	fmt.Println(animal.Name, "living in the land")
	return animal.Name+"living in the land"
}

func (animal *Animal) Info() {
	fmt.Println("the animal's name is", animal.Name, "It's", animal.Age, "years old")
}


func main() {

	//dog := new(Animal)
	//dog.Name="dog"
	//dog.Age = 12
	//dog.SpeakStyle="wangwang"

	//dog.speak()
	//dog.Info()
	//dog.breath()

	//dog :=Animal{Name: "dog", Age: 12, SpeakStyle: "wangwang"}
	//dog.speak()

	//var dog AnimalSkill
	//dog = Animal{Name: "dog", Age: 12, SpeakStyle: "wang"}
	//dog.animalspeak()

	var dog = Animal{"dog", 12, "wangwang"}
	dog.animalspeak()
	dog.breath()

}
