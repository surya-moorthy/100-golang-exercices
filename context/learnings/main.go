package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx , cancel := context.WithCancel(context.Background()) // creates a empty context.
	defer cancel()

    generators := func(dataItem string,stream chan interface{}) {
		for {
			select {
			case <-ctx.Done():
				return
			case stream <- dataItem:
			}
		}
	}	

	infiniteApples := make(chan interface{})
	go generators("apple",infiniteApples)


	infiniteOranges := make(chan interface{})
	go generators("Oranges",infiniteOranges)


	infinitePeaches := make(chan interface{})
	go generators("Peaches",infinitePeaches)

	wg.Add(1)
	go func1(ctx, &wg, infiniteApples )

	func2 := genericFunc
	func3 := genericFunc

	wg.Add(1)
	go func2(ctx, &wg, infiniteOranges )
	
	wg.Add(1)
	go func3(ctx, &wg, infinitePeaches )

	wg.Wait()
}

func func1(parentCtx context.Context, parentWg *sync.WaitGroup, stream  <-chan interface{})  {
   defer parentWg.Done()
   var wg sync.WaitGroup
   
   doWork := func(ctx context.Context) {
      defer wg.Done()
      for {
		select {
		case <-ctx.Done():
			return
		case d, ok := <-stream:
			if !ok {
				fmt.Println("channel closed...")
				return
			}
			fmt.Println(d)
		}
	  }
   }

   newCtx , cancel := context.WithTimeout(parentCtx, time.Second * 3)
   defer cancel()

   for range 3 {
	wg.Add(1)
	go doWork(newCtx)
   }

   wg.Wait()
}

func genericFunc(ctx context.Context, wg *sync.WaitGroup,stream <-chan interface{}) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case b , ok := <-stream:
			if !ok {
				fmt.Println("channel closed....")
				return
			}
			fmt.Println(b)
		}
	}
}