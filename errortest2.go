package main

import "fmt"

type divsionEle struct {
	numerator int // 分子
	denominator int //分母
}

func (div *divsionEle) error() string{
	errorMsg := `cannot proceed the function, 
				denominator is 0, you cannot use %d / 0`
	return fmt.Sprintf(errorMsg, div.numerator)
}

func divison(nmerator, denominator int) (result int,  errorMsg string){
	if denominator == 0{
		errorData := divsionEle{numerator: nmerator, denominator: denominator}
		errorMsg = errorData.error()
		return
	} else{
		return nmerator/denominator, ""
	}
}

func main() {
	result, errorMsg := divison(100, 0)
	fmt.Println(result, errorMsg)
}