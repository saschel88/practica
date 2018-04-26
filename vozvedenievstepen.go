package main

import (
	"math/big"
	"fmt"
)

func vozvedenieVStepen(x, y *big.Float) *big.Float { //Функция рекурсивного возведения безнакового целочисленного значения  в  положительную степень
	temp:=new(big.Float)
	tempSub:=new(big.Float)
	zero:= big.NewFloat(0)
	odin:=big.NewFloat(1)
	if y.Cmp(zero)==0{
		return odin
	} else if y.Cmp(odin)==0 {
		return x
	} else {
		return temp.Mul(x , vozvedenieVStepen(x, tempSub.Sub(y,odin)))// Mul
	}
}

func main() {
	x:=big.NewFloat(2)
	y:=big.NewFloat(100)
	fmt.Println(vozvedenieVStepen(x,y))


}

