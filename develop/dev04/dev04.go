// 4.	Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
)

func main() {

	// выбираем количество воркеров
	workers := 10

	// создаем канал в котрый будем писать сигнал
	quitCh := make(chan os.Signal, 1)
	// сделал так же для цикла в основном потоке
	quitCh1 := make(chan os.Signal, 1)

	// пересылаем сигналы операционной системы в канал
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	// сделал так же для цикла в основном потоке
	signal.Notify(quitCh1, syscall.SIGINT, syscall.SIGTERM)

	// создаем канал с данными
	dataCh := make(chan int)

	// создаем канал с пустой структурой для записи в него состояния
	shutdownCh := make(chan struct{})

	// создаем группу ожидания
	waitGroup := &sync.WaitGroup{}

	// создаем необходимое количество воркеров в цикле
	for i := 0; i < workers; i++ {
		// увеличиваем счетчик группы ожидания
		waitGroup.Add(1)

		// запускаем горутину в которую передаем канал, указатель на группу ожидания и порядковый номер воркера
		go func(data chan int, shutdown chan struct{}, wg *sync.WaitGroup, i int) {
			// выводим сообщение в консоль
			log.Println("Starting goroutine: ", i)
			// откладываем запуск функции декремента счетчика группы ожидания
			defer wg.Done()

			// в бесконечном цикле
			for {
				// в операторе case
				select {
				// читаем канал
				case <-shutdown:
					// если по этому каналу получены данные, то выводим в консоль сообщение
					log.Println("shutdown goroutine: ", i)
					// и выходим из цкла поcле этого горутина останавливается и счетчик группы ожидания уменьшается на единицу
					return
				case data := <-dataCh:
					fmt.Printf("Worker, %v read data: %v \n", i, data)
				// в кейсе по умолчанию
				default:
					// вроде как освобождает процессор, но как я пока не знаю
					runtime.Gosched()
				}
			}
		}(dataCh, shutdownCh, waitGroup, i)
	}

	// запускаем горутину
	go func() {
		// блокируем выполнение пока не прочитаем данные из канала
		<-quitCh
		// закрываем канал
		close(shutdownCh)
	}()

	// переменная для контроля цикла
	var stop bool = false

	for !stop {
		select {
		// читаем канал
		case <-quitCh1:
			// если по этому каналу получены данные, то выводим в консоль сообщение
			log.Println("Break main loop")
			stop = true
		default:
			// шлем в канал случайные числа
			dataCh <- rand.Intn(1000000)
		}
	}

	// ждем завершения выполнения работы всех воркеров в группе ожидания
	waitGroup.Wait()
}
