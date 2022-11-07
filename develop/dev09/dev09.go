// 9.	Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
package main

import (
	"fmt"
	"time"
)

func readMultipleSend(in <-chan int, out chan<- int) {
	for {
		out <- (<-in * 2)
	}
}

func readPrint(in <-chan int) {
	for {
		fmt.Printf("%d ", <-in)
	}
}

func main() {
	numbers := []int{21, 34, 56, 76, 34, 23, 12, 43, 46, 5, 6, 8, 4, 23, 43, 4, 34, 23, 4, 234, 23}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go readMultipleSend(ch1, ch2)
	go readPrint(ch2)

	for _, v := range numbers {
		ch1 <- v
		time.Sleep(100 * time.Millisecond)
	}

}
