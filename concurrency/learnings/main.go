package main

import (
	"fmt"
	"sync"
)

func reader(n int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		value, ok := <-ch

		if !ok {
			fmt.Println("Channel Closed")
			return
		} else {
			fmt.Printf("Hello from reader %d with value %d \n", n, value)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(4)

	go reader(1, ch, &wg)
	go reader(2, ch, &wg)
	go reader(3, ch, &wg)
	go reader(4, ch, &wg)

	for i := 0; i < 100; i++ {
		ch <- i
	}

	close(ch)

	wg.Wait()
}
