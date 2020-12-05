package main

import "fmt"

type itest interface {
	work()
}

type stest struct {
	A int
	itest
}

func main() {
	test := stest{}
	fmt.Println(test)
}

