package main

import (
	"fmt"
	"sync"
)

var totalNum int
var wg sync.WaitGroup

var lock sync.Mutex

func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		totalNum = totalNum + 1
		lock.Unlock()
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		totalNum = totalNum - 1
		lock.Unlock()
	}
}

func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Print(totalNum)
}
