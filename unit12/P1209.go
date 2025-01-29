package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var lock sync.RWMutex

func read() {
	defer wg.Done()
	lock.RLock()
	fmt.Println("start read date")
	time.Sleep(time.Second)
	fmt.Println("read successfuly ")
	lock.RUnlock()
}

func write() {

	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		fmt.Println("start write data")
		time.Sleep(time.Second * 10)
		fmt.Println(" write data successful ")
		lock.Unlock()
	}
}

func main() {
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go read()
	}

	go write()
	wg.Wait()

}
