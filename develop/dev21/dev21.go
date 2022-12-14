// 21.	Реализовать паттерн «адаптер» на любом примере.
// adapter design pattern application in go language
package main

import (
	"fmt"
)

// адаптирован
type Лампа struct {
	Включена bool
}

func (лампа *Лампа) ПолучитьСостояниеВключена() bool {
	return лампа.Включена
}

func (лампа *Лампа) УстановитьСостояниеВключена(включена bool) {
	лампа.Включена = включена
}

// цель к которой необходимо адаптировать
type ИсточникСвета interface {
	ЗажечьСвет()
}

// адаптер
type ВыключательАдаптер struct {
	лампа Лампа
}

func (выключательАдаптер *ВыключательАдаптер) ЗажечьСвет() {
	лампа := выключательАдаптер.лампа

	if лампа.ПолучитьСостояниеВключена() {
		fmt.Println("Свет включен, ничего не нужно делать")
	} else {
		лампа.УстановитьСостояниеВключена(true)
		fmt.Println("Свет выключен, включаем его")
	}
}

// клиент
func main() {
	var выключатель ИсточникСвета = &ВыключательАдаптер{
		лампа: Лампа{true},
	}

	выключатель.ЗажечьСвет()
}
