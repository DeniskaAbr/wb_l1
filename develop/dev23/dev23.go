// 23.	Удалить i-ый элемент из слайса.
package main

import "fmt"

// удалить элемент из слайса можно только созданием нового слайса

// Быстрый способ (ломает сортировку)

func DeleteElementF(a []string, i int32) {
	// Remove the element at index i from a.
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = ""   // Erase last element (write zero value).
	a = a[:len(a)-1]   // Truncate slice.
}

// Медленный способ (сохраняет сортировку)
func DeleteElementS(a []string, i int32) {
	// Remove the element at index i from a.
	copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
	a[len(a)-1] = ""     // Erase last element (write zero value).
	a = a[:len(a)-1]     // Truncate slice.
}

func main() {

	a := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	fmt.Println(a)
	var i int32 = 2

	DeleteElementF(a, i)

	fmt.Println(a)

	i = 5

	DeleteElementS(a, i)

	fmt.Println(a)
}
