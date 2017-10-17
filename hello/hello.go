package main

import (
	"test"
	"test/test1"
	"test/test2"
	"fmt"
)

func sayHello() {
	fmt.Print("say hello");
	fmt.Print("say hello ok\n");
}
func main() {
	test.RunTest();
	test11.RunTest1()
	test2.RunTest2()
	sayHello()
}
