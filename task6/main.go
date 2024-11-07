package main

import (
	"fmt"
	"time"
)

/*
Таймаут ожидания через select
Описание: Напишите программу, которая запускает горутину, выполняющую длительную операцию,
и использует select для ожидания завершения горутины или срабатывания таймаута.
Комментарий: Используйте time.After() в select, чтобы задавать время ожидания и завершить операцию, если горутина не завершилась вовремя.
*/

func main() {
	//resultChannel := make(chan int)
	exit := time.After(2 * time.Second)

	//go func() {
	//	defer close(resultChannel)
	//	time.Sleep(1 * time.Second)
	//	resultChannel <- 777
	//}()

	worker := func() <-chan int {
		resultOutChan := make(chan int)
		go func() {
			defer close(resultOutChan)
			time.Sleep(3 * time.Second)
			resultOutChan <- 888
		}()
		return resultOutChan
	}

	select {
	//case resultValue := <-resultChannel:
	//	fmt.Printf("Получившийся результат равен %v\n", resultValue)
	case resultValue := <-worker():
		fmt.Printf("Получившийся результат равен %v\n", resultValue)
	case <-exit:
		fmt.Println("Время выполнения горутины превысило допустимое время")
	}

}
