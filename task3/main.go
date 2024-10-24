package main

/*
Ожидание завершения нескольких горутин с помощью sync.WaitGroup
Описание: Напишите программу, которая запускает несколько горутин и использует sync.WaitGroup для ожидания их завершения.
Комментарий: Для каждой горутины увеличивайте счетчик через Add(),
*/

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int, 5)
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {

		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- i * i
		}()

		fmt.Println(<-ch)
	}

	wg.Wait()
	close(ch)

	for val := range ch {
		fmt.Println(val)
	}
}
