package leetcode

import "fmt"

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	lastSum := nums[0]
	curSum := 0
	curNum := nums[0]
	numsLen := len(nums)

	for i := 1; i < numsLen; i ++ {
		curNum = nums[i]
		curSum = lastSum + curNum
		if curSum >= curNum {
			lastSum = curSum
		} else {
			lastSum = curNum
		}
		if lastSum > maxSum {
			maxSum = lastSum
		}
	}

	return maxSum
}

func MaxSubArray() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	ret := maxSubArray(nums)
	fmt.Println(ret)
}