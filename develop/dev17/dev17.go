// 17.	Реализовать бинарный поиск встроенными методами языка.

package main

import "fmt"

func main() {
	// поиск выполняется по отсортированным данным
	searchField := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}

	// для каждого элемента массива выполняем поиск его значения в массиве
	for _, searchNumber := range searchField {
		fmt.Printf("Searching for %d in list %v\n", searchNumber, searchField)
		result, searchCount := binarySearch(searchField, searchNumber)
		fmt.Printf("Your number was found in position %d after %d steps.\n\n", result, searchCount)
	}
}

func binarySearch(a []int, search int) (result int, searchCount int) {
	// берем медиану массива
	mid := len(a) / 2

	switch {
	// если массив пустой то значит искать негде, ничего не найденно
	case len(a) == 0:
		result = -1 // not found
	// если искомое число больше чем значение медианы
	case a[mid] > search:
		// ищем рекурсивно в левой части массива
		result, searchCount = binarySearch(a[:mid], search)
	// если искомое число меньше чем значение медианы
	case a[mid] < search:
		// ищем рекурсивно в правой части массива
		result, searchCount = binarySearch(a[mid+1:], search)
		if result >= 0 { // if anything but the -1 "not found" result
			result += mid + 1
		}
	// если искомое число равно значению медианы
	default: // a[mid] == search
		result = mid // found
	}
	// подсчитываем количество реекурсивных вызовов
	searchCount++
	return
}
