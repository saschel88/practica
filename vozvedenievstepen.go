package main

import (
	"fmt"
)

/*func vozvedenieVStepen(x, y *big.Int) *big.Int { //Функция рекурсивного возведения безнакового целочисленного значения  в  положительную степень
	temp:=new(big.Int)
	tempSub:=new(big.Int)
	z:= big.NewInt(0)
	odin:=big.NewInt(1)
	if y ==z{
		return odin
	} else if y == odin {
		return x
	} else {
		temp.Mul(x , vozvedenieVStepen(x, tempSub.Sub(y,odin)))
		return temp
	}
}*/
func vozvedenieVStepen(x, y int64) int64 { //Функция рекурсивного возведения безнакового целочисленного значения  в  положительную степень

	if y ==0{
		return 1
	} else if y == 1 {
		return x
	} else {

		return (x * vozvedenieVStepen(x, y-1))
	}
}

func main() {

	fmt.Println(vozvedenieVStepen(2, 50))

}

