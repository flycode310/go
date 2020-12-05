package main

import (
	"fmt"
	"sync"
)

func main()  {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println(<-ch)
	}(&wg)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch)
	}()

	ch <- 1
	ch <- 1
	wg.Wait()
}
