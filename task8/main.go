package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Пул воркеров с передачей данных через канал
Описание: Реализуйте пул воркеров (горутины),
которые получают задачи из канала и обрабатывают их параллельно.
Комментарий: Создайте несколько горутин, каждая из которых будет получать задачи из одного канала и обрабатывать их.
После завершения всех задач закройте канал.
*/

func main() {
	quantityTasks := 12

	wg := &sync.WaitGroup{}

	ChannelMain := make(chan int)

	worker := func(id int) {
		defer wg.Done()

		for value := range ChannelMain {
			time.Sleep(1 * time.Second)
			fmt.Printf("Горутина %v выполнила работу - %v\n", id, value+100)
		}

	}

	go func() {

		defer close(ChannelMain)
		for i := 0; i < quantityTasks; i++ {
			ChannelMain <- i
		}
	}()

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i)
	}

	wg.Wait()

}
