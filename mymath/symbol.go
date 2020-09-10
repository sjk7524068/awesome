package mymath

func DividerPlus(x string, y int)  string{
	// 输入一个字符串x，返回一个倍数为y的字符串值
	var newX string
	for i := 0; i<y;i++{
		newX += x
	}
	return newX
}