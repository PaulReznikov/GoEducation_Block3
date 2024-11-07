package main

import (
	"fmt"
)

/*
Запуск нескольких горутин и вывод их результата в канал
Описание: Запустите несколько горутин (например, 5) и передайте результат каждой горутины в канал.
В главной функции получите значения из канала и выведите их на экран.
Комментарий: Используйте make(chan int) для создания канала и оператор go для запуска горутин.
Получение данных организуйте в цикле for.
*/

func main() {

	ch := make(chan int) // ch := make(chan int, 5)
	defer close(ch)
	//wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {

		i := i
		//wg.Add(1)
		go func() {
			//defer wg.Done()
			ch <- i * i
		}()

		fmt.Println(<-ch)
	}
	//time.Sleep(3 * time.Second)
	//wg.Wait()
	//close(ch)

	//for val := range ch {
	//	fmt.Println(val)
	//}

}
