package leetcode

import "strconv"

func summaryRanges(nums []int) []string {
	start, end, pre := 0, 0, 0
	ret := []string{}
	if len(nums) == 0 {
		return ret
	}
	start, pre = nums[0], nums[0]
	for _, v := range nums {
		if pre == v || pre +1 == v {
			pre = v
		} else {
			end = pre
			ret = append(ret, formateStr(start, end))
			start, pre = v, v
		}
	}
	end = pre
	ret = append(ret, formateStr(start, end))
	return ret
}

func formateStr(start int, end int) string {
	if start == end {
		return strconv.Itoa(start)
	} else {
		return strconv.Itoa(start) + "->" + strconv.Itoa(end)
	}
}
