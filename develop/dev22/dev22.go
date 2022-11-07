// 22.	Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

// с большими числами, которые превышают размерность стандартных типов, позволяет работать пакет 'math/big'
func main() {
	var a = new(big.Int)
	var b = new(big.Int)

	a.SetString("23456543452567543256345347865435423454213412323", 10)
	b.SetString("2345654345256754325634534786543542345", 10)

	var multi = new(big.Int)
	var div = new(big.Int)
	var sum = new(big.Int)
	var sub = new(big.Int)

	fmt.Println("%v", multi.Mul(a, b))
	fmt.Println("%v", div.Div(a, b))
	fmt.Println("%v", sum.Add(a, b))
	fmt.Println("%v", sub.Sub(a, b))
}
