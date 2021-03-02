package leetcode

import (
	"fmt"
	"testing"
)

func TestSummaryRanges(t *testing.T)  {
	nums := []int{0,1,2,4,5,7}
	ret := summaryRanges(nums)
	fmt.Println(ret)
}
