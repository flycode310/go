package leetcode

import "fmt"

func lengthOfLIS(nums []int) int {
	subsequence := []int{}
	for _, v := range nums {
		subsequence = appendItem(subsequence, v)
	}
	return len(subsequence)
}

func appendItem(subsequence []int, value int) []int{
	sLen := len(subsequence)
	if (sLen == 0) ||(subsequence[sLen-1] < value) {
		subsequence = append(subsequence, value)
		return subsequence
	}

	lowIndex := 0
	highIndex := sLen - 1
	for lowIndex <= highIndex {
		mid := (lowIndex + highIndex) / 2
		if subsequence[mid] > value {
			highIndex = mid - 1
		} else if subsequence[mid] < value {
			lowIndex = mid + 1
		} else {
			break
		}
	}

	if lowIndex > highIndex {
		subsequence[lowIndex] = value
	}
	return subsequence
}

func LengthOfLTS() {
	nums := []int{10,9,2,5,3,7,101,18}
	ret := lengthOfLIS(nums)
	fmt.Println(ret)
}