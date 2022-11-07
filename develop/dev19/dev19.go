// 19.	Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import "fmt"

func main() {

	fmt.Print("Enter text : ")
	var text string
	fmt.Scanln(&text)

	fmt.Println(reverse(text))
}

func reverse(s string) string {

	r := []rune(s)
	l := len(r)

	for i := 0; i < l/2; i++ {
		// меняем элементы местами
		r[i], r[l-1-i] = r[l-1-i], r[i]
	}

	return string(r)
}
