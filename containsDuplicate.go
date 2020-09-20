package main

func containsDuplicate(nums []int) bool {
	numMap := map[int]int{}
	for _, value := range nums {
		if _, ok := numMap[value]; ok {
			return true
		}
		numMap[value] = 1
	}
	return false
}
