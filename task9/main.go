package main

import "fmt"

/*
Конвейер обработки данных (pipeline)
Описание: Напишите программу, которая реализует конвейер обработки данных: данные проходят через несколько этапов, где на каждом этапе данные обрабатываются горутиной.
Описание задачи: Создадим конвейер, который принимает набор чисел, последовательно обрабатывает их в три этапа:
Увеличение каждого числа на 1.
Умножение каждого числа на 2.
Преобразование числа в строку с информацией о результатах обработки. (fmt.Sprintf("Initial value: %d -> Incremented: %d -> Multiplied: %d")
Комментарий: Разделите обработку данных на несколько функций, каждая из которых принимает канал для входных данных и отправляет результат в следующий канал.
*/

func IncrementValues(initCh <-chan int) <-chan int {
	incrementCh := make(chan int)

	go func() {
		defer close(incrementCh)
		for inputVal := range initCh {
			incrementCh <- inputVal + 1
		}
	}()

	return incrementCh
}

func MultiplyValues(incrementCh <-chan int) <-chan int {
	multiplyCh := make(chan int)

	go func() {
		defer close(multiplyCh)
		for incrementValue := range incrementCh {
			multiplyCh <- incrementValue * 2
		}

	}()

	return multiplyCh
}

func ConvertValuesToString(multiplyCh <-chan int) <-chan string {
	convertTostringCh := make(chan string)

	go func() {
		defer close(convertTostringCh)
		for multiplyValue := range multiplyCh {
			convertTostringCh <- fmt.Sprintf("Initial value: %d -> Incremented: %d -> Multiplied: %d", multiplyValue/2-1, multiplyValue/2, multiplyValue)
		}

	}()

	return convertTostringCh
}

func main() {
	inputSlice := []int{1, 2, 3, 4, 5, 6, 7}
	initCh := make(chan int)

	go func() {
		defer close(initCh)
		for _, value := range inputSlice {
			initCh <- value
		}
	}()

	incrementCh := IncrementValues(initCh)
	multiplyCh := MultiplyValues(incrementCh)
	convertToStringCh := ConvertValuesToString(multiplyCh)

	for resultValues := range convertToStringCh {
		fmt.Println(resultValues)
	}

}
