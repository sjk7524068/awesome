package main

import "fmt"

type man struct {
	name string
	age int
	hobby skill

}

type skill struct {
	mainHobby string
}

func (m *man) Info(){
	fmt.Println(m.name, "is", m.age)
}

func (m *skill) Info(){
	fmt.Println("hobby is", m.mainHobby)
}


func main() {
	 var woman man
	 woman.name = "lijing"
	 woman.age = 18
	 woman.hobby.mainHobby = "swiming"
	 woman.Info()
	 woman.hobby.Info()

}