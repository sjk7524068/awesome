package mymath

func IsNumBetweenRules(num int)  bool{
	if 100 > num && num >20{
		println("在区间内")
		return true
	} else{
		println("不在区间内")
		return false
	}
}
