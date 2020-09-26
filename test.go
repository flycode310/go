package main

import "fmt"

func main() {
	testS := struct{
		A string
		B string
	}{
		"A",
		"b",
	}
	testS.A = "C"
	fmt.Println(testS)
}
