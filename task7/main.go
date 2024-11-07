package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Одновременная обработка запросов с лимитом параллелизма
Описание: Напишите программу, которая запускает несколько горутин для обработки данных,
но с ограничением на количество одновременно выполняемых задач.
Комментарий: Используйте семафорный паттерн с помощью буферизованного канала или sync.WaitGroup для ограничения параллелизма.
*/

const MaxNumberGorout = 4

func main() {

	GoroutQuantRegularCh := make(chan struct{}, MaxNumberGorout)
	wg := &sync.WaitGroup{}

	worker := func(id int) {
		defer wg.Done()
		defer func() {
			<-GoroutQuantRegularCh
		}()
		GoroutQuantRegularCh <- struct{}{}

		fmt.Printf("Горутина %v начала работу\n", id)
		time.Sleep(2 * time.Second)
		//fmt.Printf("Горутина %v закончила работу\n", id)

		//<-GoroutQuantRegularCh
	}

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go worker(i)
	}

	wg.Wait()

}
