package leetcode

import (
	"fmt"
	"testing"
)

func Test697(t *testing.T) {
	nums := []int{1,2,2,3,1,4,2}
	ant := findShortestSubArray(nums)
	fmt.Println(ant)
}
