package main

import "fmt"

func main() {
	iArr := []int{5,2,7,5,2,3,5,2,3,5,2,7}
	iRetArr := []int{}

	lenArr := len(iArr)
	for i:=0; i<lenArr; i++ {
		iRetArr = append(iRetArr, 0)
	}

	for i:=1; i<lenArr; i++ {
		for j:=iRetArr[i-1]; j>=0; j-- {
			if iArr[j] == iArr[i] {
				iRetArr[i] = j + 1
				break
			}
		}
	}

	fmt.Println(iRetArr)

}
