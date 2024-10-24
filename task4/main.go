package main

import (
	"fmt"
	"sync"
)

/*
Пинг-понг между двумя горутинами через канал
Описание: Напишите две горутины, которые "перекидывают" значение через канал.
Одна горутина отправляет в канал, другая получает и снова отправляет значение.
Комментарий: Используйте канал для передачи данных между горутинами и внимательно следите за возможными дедлоками.
*/

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	num := 0

	wg.Add(2)

	go func() {
		defer wg.Done()

		for {
			ch <- num
			num = <-ch
			fmt.Println(num)
			if num > 10 {
				break
			}
		}

	}()

	go func() {
		defer wg.Done()

		for {
			num = <-ch
			ch <- num + 1
			if num == 10 {
				break
			}
		}

	}()

	wg.Wait()
}
