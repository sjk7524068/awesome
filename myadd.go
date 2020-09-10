package main

import (
	"fmt"
)

func add(x, y int)  int{
	return x+y
}

//var book map[string]string

type Speak interface {
	speak()
}

type People struct {

}

type Dog struct {

}

func (people People) speak() {
	fmt.Println("this is people speaking")
}

func (dog Dog)  speak(){
	fmt.Println("this is dog speaking")
}

func main() {
	//for i:=1; i<=10;i++{
	//	fmt.Print(mymath.Fiboo(i), " ")
	//}

	// 定义两个结构体继承于speak

	var humun Speak
	humun = new(People)
	humun.speak()

	//var humen mymath.Action
	//humen = new(mymath.PeoPle)
	//humen.Swimming()

	//var woman mymath.Man
	//woman = new(mymath.Woman)
	//fmt.Println(woman.Name(), woman.Age())
}

