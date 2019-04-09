package main

import (
	"fmt"
	"sort"
)
func main() {
	checkList := []int64{4, 1, 3, 2}
	sort.Slice(checkList, func(i, j int) bool {return checkList[i] >= checkList[j]})

	fmt.Println(checkList)
}
