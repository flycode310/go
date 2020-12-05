package main

import (
	"fmt"
	"math/rand"
	"time"
)

func request(hostname string) (response string) {
	time.Sleep(time.Duration(rand.Intn(15)) * time.Microsecond)
	return hostname
}

func main()  {
	rand.Seed(time.Now().UnixNano() / 1e6)
	responses := make(chan string, 3)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	for {
		select {
		case v :=<-responses:
			fmt.Println(v)
		case <- time.After(time.Millisecond):
			fmt.Println("return")
			return
		}
	}
}