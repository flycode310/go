package main

import (
	"fmt"
	"sync"
	"runtime"
)

var wg sync.WaitGroup
func main()  {

	runtime.GOMAXPROCS(1)
	wg.Add(1)
	chanVal := make(chan int)
	go producer(chanVal)
	consumerNum := 3
	wg.Add(consumerNum)
	for num:=1; num <= consumerNum; num++ {
		go consumer(chanVal, num)
	}

	wg.Wait()

}

func producer(in chan int) {

	for i:=0; i<10 * 100000; i++ {
		in <- i
		//fmt.Printf("g product %d\n", i)
	}
	close(in)
	wg.Done()
}

func consumer(out chan int, num int) {
	sum := 1
	for val := range out {
		sum = sumFunc(sum , val)
	}
	fmt.Printf("g%d done %d\n", num, sum)
	wg.Done()
}

func sumFunc(a int, b int) int {
	return a + b
}