// 11.	Реализовать пересечение двух неупорядоченных множеств.
package main

import (
	"fmt"
)

func Intersection(a, b []int) (c []int) {
	// для результата создаем мапу
	m := make(map[int]bool)

	// для каждого элемента массива 'а'
	for _, item := range a {
		// записываем элемент как ключ со значением bool(роли не имеет, можно и struct{})
		m[item] = true
	}

	// для массива 'b'
	for _, item := range b {
		// пытаемся записать значение как ключ в ту же мапу
		if _, ok := m[item]; ok {
			// если записать не получилось, значит на этом значении элемента произошло пересечение
			// значения на которых возникло пересечение пишем в отдельный массив
			c = append(c, item)
		}
	}
	return
}

func main() {
	// создаем два множества
	var a = []int{1, 2, 3, 4, 5, 6, 11}
	var b = []int{2, 3, 5, 6, 7}

	// выводим результат пересечения
	fmt.Println(Intersection(a, b))
}
