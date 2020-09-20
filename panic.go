package main

import (
	"fmt"
	"time"
	"net/http"
)

type IHello interface {
	SayHello() string
}

type HelloStr string

func (str HelloStr) SayHello() string {
	return string(str)
}

func PrintHello(hello IHello) {
	println(hello.SayHello())
}

func main() {

	defer func(){

		if err := recover(); err != nil {
			fmt.Printf("panic: %v", err)
		}
	}()

	defer func() {
		fmt.Printf("return message")
	}()


	http.HandleFunc("/testgo", httpTest)
	http.ListenAndServe("127.0.0.1:9999", nil)
	return
}

func httpTest(w http.ResponseWriter, req *http.Request) {

	defer func() {
		fmt.Println("defer http test return")
	}()
	fmt.Println("begin http")
	go panicTest()

	fmt.Println("return http")
	return
}
func panicTest()  {

	fmt.Println("begin sleep")
	time.Sleep(time.Duration(100 * time.Millisecond))
	fmt.Println("end sleep")

	fmt.Println("begin panic")
	//panic("panic message")
	fmt.Println("end panic")

}

