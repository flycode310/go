package main

import (
	"time"
	"fmt"
)

func main() {
	go spinner(1000 * time.Millisecond)
	const N = 45
	fibN := fib(N)
	fmt.Printf("\r Fibonacci (%d) = %d\n", N, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c",r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x - 1) + fib(x - 2)
}
