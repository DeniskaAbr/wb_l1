// 1.	Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

package main

import "fmt"

// Объявляем новый тип
type Human struct {
	name  string
	age   int
	phone string
}

// Объявляем новый тип
type Action struct {
	// встраиваем анонимное поле типа 'Human'
	Human
}

// функция для типа 'Human'
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func main() {
	// объявляем переменную типа структуры 'Human', и инициализируем поля значениями
	ivan := Human{name: "Ivan", age: 36, phone: "8(986)5-45-263"}

	// объявляем переменную типа структуры 'Action', полю 'Human' присваиваем значение
	action := Action{Human: ivan}

	// при вызове метода 'SayHi()' вызавается метод анонимного поля 'Human'
	action.SayHi()

}
