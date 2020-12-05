package main

import "fmt"

func main()  {

	sliceTest := []int{}
	go func() {
		for i:=1; i<50; i++ {
			sliceTest = append(sliceTest, i)
		}
		fmt.Println(sliceTest)
	}()
	for i:=1; i<50; i++ {
		sliceTest = append(sliceTest, i+10)
	}
	fmt.Println(sliceTest)

}
