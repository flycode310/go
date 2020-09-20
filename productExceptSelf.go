package main

func productExceptSelf(nums []int) []int {
	ret := make([]int, len(nums))
	ret[0] = 1
	for i:=1; i<len(nums); i++ {
		ret[i] = ret[i-1] * nums[i-1]
	}
	productNum := 1
	for i:=len(nums)-1; i>=0; i-- {
		ret[i] = ret[i] * productNum
		productNum *= nums[i]
	}

	return ret
}
