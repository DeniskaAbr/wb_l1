// 6.	Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"time"
)

// 1. остановка горутины через закрытие канала
// Routine 1
func channelStop() {

	ch := make(chan struct{})

	go func(ch chan struct{}) {
		for {
			_, ok := <-ch
			if !ok {
				fmt.Println("Routine 1: Stopped")
				return
			}
			fmt.Println("Routine 1: Make payload")
			// time.Sleep(time.Second)
		}
	}(ch)

	for i := 0; i < 10000000; i++ {
		if i > 200 {
			break
		} else {
			ch <- struct{}{}
		}
	}

	close(ch)
	time.Sleep(2 * time.Second)

}

// 2. останвка горутины через периодический опрос
// Routine 2
func semaphoreStop() {

	ch := make(chan struct{})
	done := make(chan struct{})

	go func() {
		for {
			select {
			case ch <- struct{}{}:
			case <-done:
				close(ch)
				return
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		done <- struct{}{}
	}()

	for _ = range ch {
		fmt.Println("Routine 2: Make payload")
	}

	fmt.Println("Routine 2: Stopped")
}

// 3. остановка горутины через различное использование контекста
// Routine 3
func contextStop() {

	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				ch <- struct{}{}
				return
			default:
				fmt.Println("Routine 3: Make payload")
			}

			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	<-ch
	fmt.Println("Routine 3: Stopped")
}

func main() {

	// запуск функции с примером остановки горутины через остановку канала
	channelStop()

	// запуск функции с примером остановки горутины через запись в канал
	semaphoreStop()

	// запуск функции с примером остановки горутины через отмену контекста
	contextStop()
}

// так же смотри задания 4 и 5
