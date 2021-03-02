package leetcode
func longestSubarray(nums []int, limit int) int {
	if len(nums) == 0 {
		return 0
	}
	ant := 0
	min, max := nums[0], nums[0]
	minIndex, maxIndex, nextIndex := 0, 0, 0
	beginIndex, endIndex := 0, 0
	for ;endIndex<len(nums);{
		if nums[endIndex] < max && nums[endIndex] > min {
			if endIndex - beginIndex + 1 > ant {
				ant = endIndex - beginIndex + 1
			}
			endIndex ++
			continue
		}

		if nums[endIndex] >= max {
			max = nums[endIndex]
			maxIndex = endIndex
			nextIndex = minIndex+1
		}
		if nums[endIndex] <= min {
			min = nums[endIndex]
			minIndex = endIndex
			nextIndex = maxIndex + 1
		}
		subValue := max - min
		if subValue <= limit {
			if endIndex - beginIndex + 1 > ant {
				ant = endIndex - beginIndex + 1
			}
			endIndex ++
		} else {
			beginIndex = nextIndex
			maxIndex, minIndex = beginIndex, beginIndex
			max, min = nums[beginIndex], nums[beginIndex]
			for j:=beginIndex; j<=endIndex; j++ {
				if nums[j] >= max {
					max = nums[j]
					maxIndex = j
				}
				if nums[j] <= min {
					min = nums[j]
					minIndex = j
				}
			}
		}
	}
	return ant
}

