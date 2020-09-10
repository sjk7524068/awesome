package mymath

import "fmt"

func Jump()  int{
	var a int = 10
	LOOP: for a <=10 && a >0{
		if a== 8{
			a=a-1
			goto LOOP
		}
		fmt.Println("now is", a)
		a--
	}

	return 0
}

func Get9x()  {
	maxN := 1
	for m:=1; m<10; m++{
		for n:=1; n<=maxN; n++{
			fmt.Printf("%dx%d=%d %v", m, n, m*n, " ")
		}
		fmt.Println()
		maxN++
	}
}

// by using goto to make get9x

func Get9xByGoTo()  {
	for m:=1; m<10; m++{
		n:=1
		LOOP: if n<=m{
			fmt.Printf("%dx%d=%d %v", m, n, m*n, " ")
			n++
			goto LOOP
		}
		fmt.Println()
		n++
	}
}

func WhoIsMax(param []int) {
	for i:=0; i< len(param);i++{
		println(param[i])
	}
}

type Books struct {
	Title string
	Author string
	BookId int
	GetTitle func()
}