package main

import (
	"time"
	"fmt"
)

func main() {

	done := make(chan struct{})

	go func() {
		time.Sleep(time.Second)
		//done <- struct{} {}
		//close(done)
		d := make(chan struct{})
		done = d
	} ()

	<- done
	fmt.Println("hello")
}
