package main

import (
	"fmt"
	"sync"
)

type SafeMapRW struct {
	mu      sync.RWMutex
	storage map[int]int
}

//func NewSafeMapRW() *SafeMapRW {
//	return &SafeMapRW{
//		storage: make(map[int]int),
//	}
//}

func (sm *SafeMapRW) Set(key, value int, wg *sync.WaitGroup) {
	defer wg.Done()
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.storage[key] = value
	fmt.Printf("Запись в мапу: [%v] = %v\n", key, value)
}

func (sm *SafeMapRW) Get(key int, wg *sync.WaitGroup) (value int, ok bool) {
	defer wg.Done()
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, ok = sm.storage[key]
	fmt.Printf("Чтение из мапы: [%v] = value: %v, ok: %v\n", key, value, ok)
	return value, ok
}

func main() {

	wg := &sync.WaitGroup{}

	taskMap := &SafeMapRW{
		storage: map[int]int{1: 10, 2: 20, 3: 30, 4: 40},
	}

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go taskMap.Get(i, wg)
	}

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go taskMap.Set(i, i*100, wg)
	}

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go taskMap.Get(i, wg)
	}

	wg.Wait()

}
