package main

import (
	"fmt"
	"sync"
)

var (
	counter       = 0
	lock          sync.Mutex
	atomicCounter = AtomicInt{}
)

type AtomicInt struct {
	value int
	lock  sync.Mutex
}

func (i *AtomicInt) Increase() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value++
}

func (i *AtomicInt) Decrease() {
	i.lock.Lock()
	defer i.lock.Unlock()

	i.value--
}

func (i *AtomicInt) Value() int {
	return i.value
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go updateCounter(&wg)
	}

	wg.Wait()
	fmt.Println(fmt.Sprintf("final counter: %d", counter))
	fmt.Println(fmt.Sprintf("final counter counter value: %d", atomicCounter.Value()))
}

func updateCounter(wg *sync.WaitGroup) {
	lock.Lock()
	defer lock.Unlock()

	counter++

	atomicCounter.Increase()
	wg.Done()
}
