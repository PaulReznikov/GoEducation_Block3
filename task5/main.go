package main

import (
	"fmt"
	"sync"
)

/*
Фан-ин и фан-аут каналов
Описание: Напишите программу, которая принимает данные из нескольких горутин (фан-ин) через один канал
и отправляет эти данные в несколько других горутин (фан-аут) для дальнейшей обработки.
Комментарий: Используйте один канал для объединения данных из нескольких источников и распределите их по другим каналам
для обработки в отдельных горутинах.
*/

func FanIn(sliceOfChannelsForFanIn ...<-chan int) <-chan int {
	outFanIn := make(chan int)

	wg := &sync.WaitGroup{}

	for _, outWorkerForFanIn := range sliceOfChannelsForFanIn {
		wg.Add(1)
		go func(chWorkerForFanIn <-chan int) {
			defer wg.Done()
			for value := range chWorkerForFanIn {
				outFanIn <- value + 100
			}
			//outFanIn <- <-chWorkerForFanIn + 100
		}(outWorkerForFanIn)
	}

	go func() {
		wg.Wait()
		defer close(outFanIn)

	}()

	return outFanIn

}

func FanOut(outFanIn <-chan int, n int) []<-chan int {
	SliceOfChannelsFanOut := make([]<-chan int, 0, n)
	wg := &sync.WaitGroup{}

	///////////////////////////////////////////////////////////////////////////////
	writeToSliseOfChannelsFanOut := func(outFanIn <-chan int) <-chan int {
		outFanOut := make(chan int)

		go func() {
			defer close(outFanOut)
			for value := range outFanIn {
				outFanOut <- value + 1000
			}
			//outFanOut <- <-c + 1000

		}()

		return outFanOut
	}
	//////////////////////////////////////////////////////////////////////////////

	for i := 0; i < n; i++ {
		i := i
		SliceOfChannelsFanOut = append(SliceOfChannelsFanOut, make(<-chan int))

		wg.Add(1)
		go func() {
			defer wg.Done()
			SliceOfChannelsFanOut[i] = writeToSliseOfChannelsFanOut(outFanIn)
		}()

	}

	wg.Wait()
	return SliceOfChannelsFanOut

}

func main() {
	wgWriteSlice := &sync.WaitGroup{}

	inputsValues := []int{1, 2, 3, 4, 5, 6}
	sliceOfChannelsForFanIn := make([]<-chan int, 0, len(inputsValues))

	/////////////Запись в слайс каналов////////////////////////////////////////////////////////////////////////////////////

	WorkerForFanIn := func(inputValue int) <-chan int {
		outWorkerForFanIn := make(chan int)
		go func() {
			defer close(outWorkerForFanIn)
			outWorkerForFanIn <- inputValue + 10
		}()
		return outWorkerForFanIn
	}

	for i := 0; i < len(inputsValues); i++ {
		i := i
		sliceOfChannelsForFanIn = append(sliceOfChannelsForFanIn, make(<-chan int))
		wgWriteSlice.Add(1)
		go func() {
			defer wgWriteSlice.Done()
			sliceOfChannelsForFanIn[i] = WorkerForFanIn(inputsValues[i])
		}()
	}

	wgWriteSlice.Wait()
	/////////////////////////////////////////////////////////////////////////////////////////////////

	//for v := range FanIn(sliceOfChannelsForFanIn...) {
	//	fmt.Println(v)
	//}

	for _, channel := range FanOut(FanIn(sliceOfChannelsForFanIn...), len(inputsValues)) {
		fmt.Println(<-channel)
	}

}
