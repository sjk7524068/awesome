package main

import (
	"awesomeProject/mymath"
	"fmt"
)

var s1 = "google is the best"


const bee = "i'm bee"


func main()  {
	//fmt.Println("hello world~")
	//// 调用main packge 里的方法
	//fmt.Println(add(10, 600))
	//// 调用mymath包里的方法
	//var result = mymath.Add2(100, 200)
	//fmt.Println(result)
	//fmt.Println(mymath.Sub(100, 50))
	////字符串拼接
	//fmt.Println("Google"+"is the best")
	//fmt.Println(s1[3:7])
	//fmt.Println(s1[3])
	//fmt.Println(len(s1))
	//// 调用一个方法
	//strloop()
	//tList := []int{10, 50, 60, 13, 58, 57, 7, 83}

	//// 设置一个随机种子
	//rand.Seed(int64(time.Now().UnixNano()))
	////var tList [] int
	//for i:=0; i<10; i++{
	//	value := rand.Intn(100)
	//	println(&value, value)
	//	//tList = append(tList, rand.Intn(100))
	//}
	//println(mymath.DividerPlus("_______", 10))
	//// 真随机数
	//for i:=0; i<10; i++{
	//	addr, _ :=cryrand.Int(cryrand.Reader, big.NewInt(100))
	//	fmt.Println(addr, &addr)
	//}


	//var balance = [] float32{1.23, 50.65, 1.0035}
	//fmt.Println(balance, len(balance))
	//
	//var mulgroup = [2][3]int{
	//	{10, 25, 6},
	//	{5, 80, 41},
	//}
	//fmt.Println(mulgroup)

	//var group = []int{10, 52, 86, 9}
	//mymath.WhoIsMax(group)

	//var countryName map[string]string{"china":"中国"}
	countryName := make(map[string]string)
	countryName["china"] = "中国"
	countryName["USA"] = "美国"
	delete(countryName, "USA")
	fmt.Println(countryName)


}

type Books2 struct {
	title string
	author string
	bookId int
}

func strloop()  {
	// 开始utf-8遍历
	//for i:=0; i<len(s1);i++{
	//	ch := s1[i]
	//	fmt.Println(ch)
	//}
	println(mymath.DividerPlus("______", 10))
	//unicode 遍历
	for _, ch1 := range s1 {
		if "g" == string(ch1){
			fmt.Println(string(ch1))
		} else{
			println("g is not exist")
		}
	}

}
