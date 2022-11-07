// 16.	Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
package main

import "fmt"

func main() {
	fmt.Println(quickSortStart([]int{2, 50, 23, 34, 5, 2, 87, 23, 54, 7, 4, 3, 2, 343, 423, 423, 42, 5, 46, 5464}))
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		// бъем массив предварительной сортировкой на части
		arr, p = partition(arr, low, high)
		// сортируем части до позиции опорного элемента
		arr = quickSort(arr, low, p-1)
		// сортируем части после позиции опорного элемента
		arr = quickSort(arr, p+1, high)
	}
	// возвращаем отвортированный массив
	return arr
}

// старт сортировки
func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

// бъем на части
func partition(arr []int, low, high int) ([]int, int) {
	// опорный элемент
	pivot := arr[high]

	i := low
	// в цикле оббегаем по элементам
	for j := low; j < high; j++ {
		// сравниваем все осстальные элементы с опорным
		if arr[j] < pivot {
			// переставляем в массиве
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	// переставляем первый и последний элементы если необходимо
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
