/*
15.	К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string

	func someFunc() {
	  v := createHugeString(1 << 10)
	  justString = v[:100]
	}

	func main() {
	  someFunc()
	}
*/

// желательно использовать условия при создании слайса чтоб не выйти за границы
// justString это переменная типа string с фиксированной длиной, желательно не использовать глобальную переменную,
// а инициализировать и возвращать новую при необходимости
// слайс от строки берется как от массива байтов по этому для исключения разрыва с данных о символах,
// необходимо работать со строкой как с массивом рун(символов) правда это приводит к использованию
// дополнительных переменных для перобразования значений

package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var letters = []rune("ᚠᚡᚢᚣᚤᚥᚦᚧᚨᚩᚪᚫᚬᚭᚮᚯᚰᚱᚲᚳᚴᚵᚶᚷᚸᚹᚺᚻᚼᚽᚾᚿᛀᛁᛂᛃᛄᛅᛆᛇᛈᛉᛊᛋᛌᛍᛎᛏᛐᛑᛒᛓᛔᛕᛖᛗᛘᛙᛚᛛᛜᛝᛞᛟᛠᛡᛢᛣᛤᛥᛦᛧᛨᛩᛪ᛫᛬᛭ᛮᛯᛰ")

func someFunc() ([]rune, error) {
	h := 0
	l := 100

	v := *createHugeString(1 << 10)

	if h > len(v)-1 || l > len(v)-1 {
		return []rune{}, error(errors.New("Values in over bounds"))
	}

	justString := v[:100]

	return justString, nil
}

func createHugeString(l int) *[]rune {
	b := make([]rune, l)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return &b
}

func main() {
	r, err := someFunc()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(r))
}
