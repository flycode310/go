package main

import (
	"fmt"
)

var a = ""
var ch = make(chan int)
func f(strA string) {

	funcV := func() {
		fmt.Println(strA)
	}
	funcV()
	<- ch
}
func main() {

	go f("hello world 123")
	ch <- 0
	fmt.Println("finish")


}
