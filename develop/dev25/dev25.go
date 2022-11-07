// 25.	Реализовать собственную функцию sleep.

package main

import (
	"context"
	"fmt"
	"time"
)

// спит до завершения контекста
func SleepCtx(d time.Duration) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()

	if <-ctx.Done(); true {
		return
	}
}

// спит до завершения таймера
func SleepTimer(d time.Duration) {
	t := time.NewTimer(d)
	if <-t.C; true {
		return
	}
}

// спит до тех пор пока в канал не поступит время от функции After
func SleepAfter(d time.Duration) {
	if <-time.After(d); true {
		return
	}
}

// спит пока проверка текущего времени не удовлетворит условию
func SleepLoop(d time.Duration) {
	startTime := time.Now()

	for {
		end := time.Now()
		sub := end.Sub(startTime) / time.Nanosecond
		if sub >= d {
			break
		}
	}
}

func main() {

	d := time.Duration(15) * time.Second
	fmt.Printf("Текущее время: %s\n", time.Now())

	go func() {
		SleepCtx(d)
		fmt.Printf("Время после выполнения SleepCtx: %s\n", time.Now())
	}()

	go func() {
		SleepTimer(d)
		fmt.Printf("Время после выполнения SleepTimer: %s\n", time.Now())
	}()

	go func() {
		SleepAfter(d)
		fmt.Printf("Время после выполнения SleepAfter: %s\n", time.Now())
	}()

	go func() {
		SleepLoop(d)
		fmt.Printf("Время после выполнения SleepLoop: %s\n", time.Now())
	}()

	dn := time.Duration(20) * time.Second
	SleepCtx(dn)
	fmt.Println("The End!")

}
