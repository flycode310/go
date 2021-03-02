package main

import "fmt"

type testStruct struct {
	a string
	b int
}

func main() {
	t1 := testStruct{
		a: "123",
		b: 1,
	}
	fmt.Println(t1)
	t2 := t1
	t2.b = 2
	fmt.Println(t1)
	fmt.Println(t2)

	t3 := &t1
	t3.b = 3
	fmt.Println(t1)
	fmt.Println(t3)
}
