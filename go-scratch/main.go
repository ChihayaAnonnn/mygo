package main

import "sync"

var c = make(chan int, 10)
var a string

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		a = "Hello, World!"
		wg.Done()
	}()

	wg.Wait()
	print(a)
}
