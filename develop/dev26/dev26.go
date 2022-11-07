/*
26.	Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

package main

import (
	"fmt"
	"strings"
)

// функция которая проверяет перебором
func CheckIsContainsUnicalSymbolsLoop(st string) bool {
	s := strings.ToLower(st)
	var unic = true

	for i, r := range s {
		for n, r1 := range s {
			if r == r1 && i != n {
				unic = false
				break
			}
		}
		if !unic {
			break
		}
	}

	if unic {
		return unic
	} else {
		return unic
	}
}

// функция которая проверяет попыткой записи в мапу
func CheckIsContainsUnicalSymbolsPushToMap(st string) bool {
	s := strings.ToLower(st)
	var unic = true

	m := make(map[rune]struct{})
	for _, r := range s {
		_, ok := m[r]
		if !ok {
			m[r] = struct{}{}
		} else {
			unic = false
			break
		}
	}
	return unic
}

func main() {

	fmt.Printf("abcd - %v \n", CheckIsContainsUnicalSymbolsLoop("abcd"))
	fmt.Printf("abCdefAaf - %v \n", CheckIsContainsUnicalSymbolsLoop("abCdefAaf"))
	fmt.Printf("aabcd - %v \n", CheckIsContainsUnicalSymbolsLoop("aabcd"))

	fmt.Printf("abcd - %v \n", CheckIsContainsUnicalSymbolsPushToMap("abcd"))
	fmt.Printf("abCdefAaf - %v \n", CheckIsContainsUnicalSymbolsPushToMap("abCdefAaf"))
	fmt.Printf("aabcd - %v \n", CheckIsContainsUnicalSymbolsPushToMap("aabcd"))
}
