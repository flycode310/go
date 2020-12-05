package leetcode

func longestConsecutive(nums []int) int {
	numSet := make(map[int]int)
	for _, v := range nums {
		numSet[v] =  0
	}

	maxNum := 0
	for _, v := range nums {
		len := 1
		if numSet[v] > 0 {
			len = numSet[v]
		} else {
			len = getLength(v, &numSet)
		}
		if maxNum < len {
			maxNum = len
		}
	}

	return maxNum
}

func getLength(value int, numSet *map[int]int) int {
	if _, ok:= (*numSet)[value]; !ok {
		return 0
	} else if (*numSet)[value] > 0 {
		return (*numSet)[value]
	} else {
		len := getLength(value-1, numSet)
		(*numSet)[value] = len + 1
		return (*numSet)[value]
	}
}
