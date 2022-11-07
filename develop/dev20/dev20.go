// 20.	Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».
package main

import (
	"fmt"
	"strings"
)

func main() {

	var text string = "snow dog sun — sun dog snow"

	fmt.Println(reverse(text))
}

func reverse(s string) string {
	ss := strings.Split(s, " ")

	l := len(ss)
	var sb string

	for i := l - 1; i >= 0; i-- {

		if i != l-1 {
			sb = sb + " "
		}
		sb = sb + ss[i]
	}

	return sb
}
