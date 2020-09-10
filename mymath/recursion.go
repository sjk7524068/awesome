package mymath

import "fmt"

func Fiboo(x int)  int{
	if x<2{
		return x
	} else {
		return Fiboo(x-1)+Fiboo(x-2)
	}
}


type Action interface {
	Swimming()
}

type PeoPle struct {

}

func(people PeoPle) Swimming(){
	fmt.Println("humen can swimming")
}

type Man interface {
	Name() string
	Age() int
}

type Woman struct {

}

func (woman Woman)  Name() string{
	return "jesscia"
}

func (woman Woman)  Age() int{
	return 18
}

