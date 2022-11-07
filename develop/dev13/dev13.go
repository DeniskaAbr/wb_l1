// 13.	Поменять местами два числа без создания временной переменной.
package main

import "fmt"

func main() {
	firstValue := 1
	secondValue := 2

	fmt.Printf("firstValue is: %v, secondValue is: %v\n", firstValue, secondValue)

	// swap

	firstValue, secondValue = secondValue, firstValue

	fmt.Println("after swap")
	fmt.Printf("firstValue is: %v, secondValue is: %v\n", firstValue, secondValue)

}
