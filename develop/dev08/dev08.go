// 8.	Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
package main

import "fmt"

func BitReplace(n int64, p int, c bool) int64 {
	if c {
		n |= 1 << p // устанавливаем байт на позиции при помощи оператора ИЛИ
	} else {
		n &^= 1 << p // сбрасываем байт на позиции при помощи оператора И НЕ
	}
	return n
}

func main() {
	var number int64 = 50

	fmt.Printf("%d in binary %b\n", number, number)
	number = BitReplace(number, 3, true)
	fmt.Printf("%d in binary %b\n", number, number)
	number = BitReplace(number, 3, false)
	fmt.Printf("%d in binary %b\n", number, number)
	number = BitReplace(number, 2, true)
	fmt.Printf("%d in binary %b\n", number, number)
	number = BitReplace(number, 3, true)
	fmt.Printf("%d in binary %b\n", number, number)

}

// А может нужно было  менять через unsafe в памяти?
