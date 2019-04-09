package main

import (
	"sort"
	"strings"
	"fmt"
)
func main() {
	sortSlice()
	sortTest()
}

func sortSlice() {

	var checkList []int64
	checkList = append(checkList, 100)
	checkList = append(checkList, 98)
	checkList = append(checkList, 200)
	checkList = append(checkList, 1)
	sort.Slice(checkList, func(i, j int) bool {return checkList[i] <= checkList[j]})

	strCheckList := strings.Replace(strings.Trim(fmt.Sprint(checkList), "[]"), " ", ",", -1)
	fmt.Println(strCheckList)
}

func sortTest() {
	panic(2)
	data := make(map[string]string)
	if data == nil {
		fmt.Println("map nil")
	} else {
		fmt.Println("map not nil")
	}

	var checkNoList []int64
	checkNoList = append(checkNoList, 677);
	sort.Slice(checkNoList, func(i, j int) bool {
		return checkNoList[i] <= checkNoList[j]
	})
	data["check_no_list"] = strings.Replace(strings.Trim(fmt.Sprint(checkNoList), "[]"), " ", ",", -1)

	var hitNoList []int64
	sort.Slice(hitNoList, func(i, j int) bool {
		return hitNoList[i] <= hitNoList[j]
	})
	data["hit_no_list"] = strings.Replace(strings.Trim(fmt.Sprint(hitNoList), "[]"), " ", ",", -1)

	fmt.Println(data)
}
