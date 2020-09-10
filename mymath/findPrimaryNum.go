package mymath

import "fmt"


func FindPrimaryNum(IntList []int)  []int{
	// 输入一个整型数组切片，返回其宿舍集合切片
	var PMN []int
	for _, ele := range IntList{
		pcount := 0
		for count:=2; count<=ele; count++{
			if ele%count == 0{
				pcount += 1
				if pcount > 1{
					continue
				}
				//println(ele, "can % by", count)
			}
		}
		if pcount ==1 {
			println(ele, "is PrimaryNum")
			PMN = append(PMN, ele)
		}
	}
	fmt.Println(DividerPlus("_____", 10))
	fmt.Println(PMN)
	fmt.Printf("%v\n", PMN)
	printSlice(PMN)
	return PMN
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
